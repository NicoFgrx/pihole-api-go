package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	fmt.Println("[+] Get pihole status")
	status, err := client.GetStatus()
	if err != nil {
		log.Fatalf("An error occured fetching status : %s", err)
	}
	fmt.Printf("[+] Status founded : %s\n", status)

	fmt.Println("[+] Disable blocking for 10 sec")
	status, err = client.DisableBlocking(10)
	if err != nil {
		log.Fatalf("An error occured fetching status : %s", err)
	}

	fmt.Printf("[+] Status founded : %s\n", status)

	fmt.Printf("[+] Wait 11 seconds\n")
	time.Sleep(11 * time.Second)

	// Go back to Enabled status after 10 secondes
	fmt.Println("[+] Get pihole status")
	status, err = client.GetStatus()
	if err != nil {
		log.Fatalf("An error occured fetching status : %s", err)
	}
	fmt.Printf("[+] Status founded : %s\n", status)

	fmt.Println("[+] Disable blocking for unlimited time")
	status, err = client.DisableBlocking()
	if err != nil {
		log.Fatalf("An error occured fetching status : %s", err)
	}
	fmt.Printf("[+] Status founded : %s\n", status)

	fmt.Println("[+] Enable blocking")
	status, err = client.EnableBlocking()
	if err != nil {
		log.Fatalf("An error occured fetching status : %s", err)
	}
	fmt.Printf("[+] Status founded : %s\n", status)

}
