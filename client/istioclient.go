package client

import (
	"log"

	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
)

type IstioClient struct {
	Kubeconfig string
}

func (ic *IstioClient) CreateClientSet() (clientset *versionedclient.Clientset) {

	// get config from file or fallback to incluster
	config, err := clientcmd.BuildConfigFromFlags("", ic.Kubeconfig)
	if err != nil {
		log.Fatalf("Failed to create k8s client: %s", err)
	}

	clientset, err = versionedclient.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}

	return clientset
}
