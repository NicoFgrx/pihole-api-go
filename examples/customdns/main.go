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

	fmt.Println("[+] Get all existing dns records")
	customdns, err := client.GetAllCustomDNS()
	if err != nil {
		log.Fatalf("An error occured while fetching all dns records : %s", err)
	}

	for _, record := range customdns {
		fmt.Printf("	- %s: %s\n", record.Domain, record.IP)
	}

	fmt.Println("[+] Create new dns records")
	err = client.AddCustomDNS(
		&pihole.DNSRecordParams{
			Domain: "test3.example.com",
			IP:     "3.3.3.3",
		},
	)
	if err != nil {
		log.Fatalf("An error occured while create new record: %s", err)
	}

	fmt.Println("[+] Get the new dns records only")
	customdnsrecord, err := client.GetCustomDNS("test3.example.com")
	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	fmt.Printf("	- %s: %s\n", customdnsrecord.Domain, customdnsrecord.IP)

	fmt.Println("[+] Update the new dns records with different IP")
	err = client.UpdateCustomDNS("test3.example.com", &pihole.DNSRecordParams{
		Domain: "test3.example.com",
		IP:     "33.33.33.33",
	})
	if err != nil {
		log.Fatalf("An error occured while update dns record : %s", err)
	}

	fmt.Println("[+] Get the new dns records only")
	customdnsrecord, err = client.GetCustomDNS("test3.example.com")
	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}
	fmt.Printf("	- %s: %s\n", customdnsrecord.Domain, customdnsrecord.IP)

	fmt.Println("[+] Delete new dns records")
	err = client.DeleteCustomDNS(
		&pihole.DNSRecordParams{
			Domain: "test3.example.com",
			IP:     "33.33.33.33",
		},
	)

	if err != nil {
		log.Fatalf("An error occured while deleting the dns record : %s", err)
	}

	fmt.Println("[+] Get all existing dns records")
	customdns, err = client.GetAllCustomDNS()
	if err != nil {
		log.Fatalf("An error occured while fetching all dns records: %s", err)
	}

	for _, record := range customdns {
		fmt.Printf("	- %s: %s\n", record.Domain, record.IP)
	}

}
