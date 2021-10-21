package security

import (
	"context"
	"denytor/tor"
	"log"

	securityv1beta1 "istio.io/api/security/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	istiotypev1beta1 "istio.io/api/type/v1beta1"
	"istio.io/client-go/pkg/apis/security/v1beta1"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
)

const (
	AuthPolicyName = "ingressgateway-deny-tor"
	Namespace      = "istio-system"
	LabelApp       = "app"
	LabelAppValue  = "istio-ingressgateway"
)

func CreateAuthorizationPolicyV1Beta1(cs *versionedclient.Clientset, url string, remoteIpBlock bool) {

	// create new vs
	ctx := context.Background()

	torIpList := tor.GetTorExitNodesList(url)

	var from *securityv1beta1.Rule_From

	// check client Ip
	if remoteIpBlock {
		// client Ip populated from X-Forwarded-For header or proxy protocol
		from = &securityv1beta1.Rule_From{
			Source: &securityv1beta1.Source{
				RemoteIpBlocks: torIpList,
			},
		}
	} else {
		// client Ip from passthrough LB
		from = &securityv1beta1.Rule_From{
			Source: &securityv1beta1.Source{
				IpBlocks: torIpList,
			},
		}
	}

	// define deny ips rule
	rule := &securityv1beta1.Rule{
		From: []*securityv1beta1.Rule_From{from},
	}

	// devine policy with rules
	authPolicy := &v1beta1.AuthorizationPolicy{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      AuthPolicyName,
			Namespace: Namespace,
		},
		Spec: securityv1beta1.AuthorizationPolicy{
			Selector: &istiotypev1beta1.WorkloadSelector{
				MatchLabels: map[string]string{LabelApp: LabelAppValue},
			},
			Rules:  []*securityv1beta1.Rule{rule},
			Action: securityv1beta1.AuthorizationPolicy_DENY,
		},
	}

	existentAuthPolicy, err := cs.SecurityV1beta1().AuthorizationPolicies(Namespace).Get(ctx, AuthPolicyName, metav1.GetOptions{})

	// create or update existent istio AuthorizationPolicy
	if errors.IsNotFound(err) {
		_, err = cs.SecurityV1beta1().AuthorizationPolicies(Namespace).Create(ctx, authPolicy, metav1.CreateOptions{})

		if err != nil {
			log.Fatalf("Error creating AuthorizationPolicy [ %s ].", err)
		}
		log.Printf("Istio AuthorizationPolicy [ %s ] created.", AuthPolicyName)
	} else {

		authPolicy.ObjectMeta.ResourceVersion = existentAuthPolicy.ObjectMeta.ResourceVersion

		_, err = cs.SecurityV1beta1().AuthorizationPolicies(Namespace).Update(ctx, authPolicy, metav1.UpdateOptions{})

		if err != nil {
			log.Fatalf("Error updating AuthorizationPolicy [ %s ].", err)
		}
		log.Printf("Istio AuthorizationPolicy [ %s ] updated.", AuthPolicyName)
	}

}
