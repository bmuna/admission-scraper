package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type University struct {
	Name     string   `json:"name"`
	WebPages []string `json:"web_pages"`
	Country  string   `json:"country"`
}

func main() {
	fmt.Println("Hello, World!")

	resp, err := http.Get("https://raw.githubusercontent.com/Hipo/university-domains-list/refs/heads/master/world_universities_and_domains.json")

	if err != nil {
		fmt.Printf("error when fetching the api %v", err)
	}

	defer resp.Body.Close()

	var universities []University

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error when fetching the api %v", err)
	}

	err = json.Unmarshal(body, &universities)
	if err != nil {
		fmt.Printf("error unmarshaling JSON: %v", err)
		return
	}

	fmt.Printf("Loaded %d universities worldwide\n", len(universities))

}
