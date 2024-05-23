package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached the Player Graph Service")
	})

	// Fetch matches as part of the service start-up for testing
	puuid := os.Getenv("SOURCE_PUUID")
	region := os.Getenv("REGION")
	apiKey := os.Getenv("RIOT_API_KEY")

	matches := fetchMatches(puuid, region, apiKey)
	log.Println("Fetched matches: ", matches)

	fmt.Println("Player Graph Service listening on http://localhost:8083")
	http.ListenAndServe(":8083", nil)
}
