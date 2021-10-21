package main

import (
	"denytor/client"
	"denytor/security"
	"flag"
)

func main() {

	// Parse cmd flags
	kubeconfig := flag.String("kubeconfig", "", "Kubeconfig file path. Only necessary outside kubernetes cluster. (Optional)")
	remoteIpBlock := flag.Bool("remote-ip-block", true, "Client IP populated from X-Forwarded-For header or proxy protocol?")
	torURL := flag.String("tor-exit-nodes-list-url", "https://check.torproject.org/torbulkexitlist", "URL to get TOR exit nodes list. Response content type must be [ text/plain ] ")
	flag.Parse()

	ic := client.IstioClient{
		Kubeconfig: *kubeconfig,
	}

	// create Authorization Policy
	security.CreateAuthorizationPolicyV1Beta1(ic.CreateClientSet(), *torURL, *remoteIpBlock)

}
