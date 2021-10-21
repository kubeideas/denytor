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

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error accessing url [ %s ]: %s", url, err)
	}

	defer resp.Body.Close()

	if !strings.Contains(resp.Header.Get("content-type"), "text/plain") {
		log.Fatalf("Invalid content type! Expected [ text/plain; charset=utf-8 ] but received [ %s ]. Please check URL informed.", resp.Header.Get("content-type"))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error getting body: ", err)
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
