package tor

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//GetTorExitNodesList retrieves Tor exit ip list
func GetTorExitNodesList(url string) []string {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Error creating http request: %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error getting reponse for GET url [ %s ]: %s", url, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error getting body: %s", err)
	}

	// check body type
	detectedType := http.DetectContentType(body)
	if !strings.Contains(detectedType, "text/plain") {
		log.Fatalf("Invalid content type! Expected [ text/plain; charset=utf-8 ] but received [ %s ]. Please check URL informed.", detectedType)
	}

	log.Printf("Using TOR URL [ %s ] to create ip list.", url)

	// trim spaces
	ipList := strings.Split(string(bytes.TrimSpace(body)), "\n")

	// check if ipList is empty
	if len(ipList) == 0 || ipList[0] == "" {
		log.Fatalln("Ip list is empty. Check URL provided!")
	}

	log.Printf("Ip list has [ %d ] entries.", len(ipList))

	return ipList
}
