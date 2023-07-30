package main

import (
	"encoding/json"
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

	fmt.Println("[+] Get existing dns records")
	res, err := client.GetCustomDNS()
	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	defer res.Body.Close()

	var post pihole.GetCustomDNSResponse

	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	for i := 0; i < len(post.Data); i++ {
		fmt.Printf("	- %s: %s\n", post.Data[i][0], post.Data[i][1])
	}

	fmt.Println("[+] Create new dns records")
	res, err = client.AddCustomDNS(
		&pihole.DNSRecordParams{
			Domain: "box.pasfastoche.lan",
			IP:     "192.168.1.1",
		},
	)
	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	var post2 pihole.PostCustomDNSResponse

	if err := json.NewDecoder(res.Body).Decode(&post2); err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	fmt.Println(post2)

	fmt.Println("[+] Get existing dns records")
	res, err = client.GetCustomDNS()

	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	for i := 0; i < len(post.Data); i++ {
		fmt.Printf("	- %s: %s\n", post.Data[i][0], post.Data[i][1])
	}

	fmt.Println("[+] Delete new dns records")
	res, err = client.DeleteCustomDNS(
		&pihole.DNSRecordParams{
			Domain: "box.pasfastoche.lan",
			IP:     "192.168.1.1",
		},
	)

	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	if err := json.NewDecoder(res.Body).Decode(&post2); err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	fmt.Println(post2)

	fmt.Println("[+] Get existing dns records")
	res, err = client.GetCustomDNS()

	if err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	for i := 0; i < len(post.Data); i++ {
		fmt.Printf("	- %s: %s\n", post.Data[i][0], post.Data[i][1])
	}

}
