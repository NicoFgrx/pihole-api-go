package main

import (
	"fmt"
	"log"
	"os"

	pihole "github.com/NicoFgrx/pihole-api-go/api"
)

func main() {

	url := os.Getenv("PIHOLE_API_URL") // must be http[s]://<IP>:<port>/admin/api.php
	key := os.Getenv("PIHOLE_TOKEN")

	fmt.Println("[+] Creating client")

	client := pihole.NewClient(url, key)

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
			Domain: "box.pasfastoche.lan",
			IP:     "192.168.1.1",
		},
	)
	if err != nil {
		log.Fatalf("An error occured while create new record: %s", err)
	}

	fmt.Println("[+] Get the new dns records only")
	customdnsrecord, err := client.GetCustomDNS("box.pasfastoche.lan")
	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	fmt.Printf("	- %s: %s\n", customdnsrecord.Domain, customdnsrecord.IP)

	fmt.Println("[+] Update the new dns records with different IP")
	err = client.UpdateCustomDNS("box.pasfastoche.lan", &pihole.DNSRecordParams{
		Domain: "box.pasfastoche.lan",
		IP:     "192.168.1.2",
	})
	if err != nil {
		log.Fatalf("An error occured while update dns record : %s", err)
	}

	fmt.Println("[+] Get the new dns records only")
	customdnsrecord, err = client.GetCustomDNS("box.pasfastoche.lan")
	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}
	fmt.Printf("	- %s: %s\n", customdnsrecord.Domain, customdnsrecord.IP)

	fmt.Println("[+] Delete new dns records")
	err = client.DeleteCustomDNS(
		&pihole.DNSRecordParams{
			Domain: "box.pasfastoche.lan",
			IP:     "192.168.1.2",
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
