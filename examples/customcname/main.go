package main

import (
	"fmt"
	"log"
	"os"

	pihole "github.com/NicoFgrx/pihole-api-go/api"
)

func Config() pihole.Client {
	url := os.Getenv("PIHOLE_API_URL") // must be http[s]://<IP>:<port>/admin/api.php
	key := os.Getenv("PIHOLE_TOKEN")

	if url == "" {
		url = "http://localhost:8080/admin/api.php"
	}
	if key == "" {
		key = "96cf46f9e9312ea9ad00f5f9e63b25643f701246357068549a6c2ea3d163bf1e"
	}

	fmt.Println("[+] Creating client")

	client := pihole.NewClient(url, key)

	return *client
}

func main() {

	client := Config()

	fmt.Println("[+] Get all existing CNAME records")
	customCNAME, err := client.GetAllCustomCNAME()
	if err != nil {
		log.Fatalf("An error occured while fetching all CNAME records : %s", err)
	}

	for _, record := range customCNAME {
		fmt.Printf("	- %s -> %s\n", record.Domain, record.Target)
	}

	fmt.Println("[+] Create new DNS records")
	err = client.AddCustomDNS(
		&pihole.DNSRecordParams{
			Domain: "test3.example.com",
			IP:     "3.3.3.3",
		},
	)
	if err != nil {
		log.Fatalf("An error occured while create new dns record: %s", err)
	}

	fmt.Println("[+] Create new CNAME records")
	err = client.AddCustomCNAME(
		&pihole.CNAMERecordParams{
			Domain: "web",
			Target: "test3.example.com",
		},
	)
	if err != nil {
		log.Fatalf("An error occured while create new record: %s", err)
	}

	fmt.Println("[+] Get the new CNAME records only")
	customCNAMErecord, err := client.GetCustomCNAME("web")
	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	fmt.Printf("	- %s -> %s\n", customCNAMErecord.Domain, customCNAMErecord.Target)

	fmt.Println("[+] Update the new CNAME records with different alias")
	err = client.UpdateCustomCNAME("web", &pihole.CNAMERecordParams{
		Domain: "web2",
		Target: "test3.example.com",
	})
	if err != nil {
		log.Fatalf("An error occured while update CNAME record : %s", err)
	}

	fmt.Println("[+] Get the new CNAME records only")
	customCNAMErecord, err = client.GetCustomCNAME("web2")
	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}
	fmt.Printf("	- %s -> %s\n", customCNAMErecord.Domain, customCNAMErecord.Target)

	fmt.Println("[+] Delete new CNAME records")
	err = client.DeleteCustomCNAME(
		&pihole.CNAMERecordParams{
			Domain: "web2",
			Target: "test3.example.com",
		},
	)
	if err != nil {
		log.Fatalf("An error occured while deleting the CNAME record : %s", err)
	}

	fmt.Println("[+] Get all existing CNAME records")
	customCNAME, err = client.GetAllCustomCNAME()
	if err != nil {
		log.Fatalf("An error occured while fetching all CNAME records: %s", err)
	}

	for _, record := range customCNAME {
		fmt.Printf("	- %s -> %s\n", record.Domain, record.Target)
	}

}
