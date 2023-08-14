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

	fmt.Println("[+] Get all white domain")
	custom_cat := "white" // can be white, black, regex_white, regex_black
	managed_domains_lst, err := client.GetManagedDomainsFromCategory(custom_cat)
	if err != nil {
		log.Fatalf("An error occured %s fetching all white domain : %s", custom_cat, err)
	}

	for _, record := range managed_domains_lst {
		fmt.Printf("	- %s\n", record.Domain)
	}

	fmt.Println("[+] Get all invalid-tested-category domain")
	managed_domains_lst, err = client.GetManagedDomainsFromCategory("invalid-tested-category")
	if err != nil {
		fmt.Printf("(skipped) An error occured while fetching all invalid-tested-category domain : %s\n", err)
		// log.Fatalf("An error occured while fetching all invalid-tested-category domain : %s", err)
	}

	for _, managed_domain := range managed_domains_lst {
		fmt.Printf("	- %s\n", managed_domain.Domain)
	}

	fmt.Println("[+] Create example managed domain (domain : white, test1.example.com)")
	err = client.AddManagedDomains("white", "test1.example.com")
	if err != nil {
		log.Fatalf("An error occured while creating example managed domain : %s", err)
	}

	fmt.Println("[+] Get all managed domains")
	managed_domains_lst, err = client.GetAllManagedDomains()
	if err != nil {
		log.Fatalf("An error occured while fetching all managed domain : %s", err)
	}

	for _, managed_domain := range managed_domains_lst {
		fmt.Printf("	- {category : %s, domain : %s}\n", pihole.Categories[managed_domain.Type], managed_domain.Domain)
	}

	fmt.Println("[+] Delete example managed domain (domain : white, test1.example.com)")
	err = client.SubManagedDomains("white", "test1.example.com")
	if err != nil {
		log.Fatalf("An error occured while creating example managed domain : %s", err)
	}

	fmt.Println("[+] Get all managed domains")
	managed_domains_lst, err = client.GetAllManagedDomains()
	if err != nil {
		log.Fatalf("An error occured while fetching all managed domain : %s", err)
	}

	for _, managed_domain := range managed_domains_lst {
		fmt.Printf("	- {category : %s, domain : %s}\n", pihole.Categories[managed_domain.Type], managed_domain.Domain)
	}

}
