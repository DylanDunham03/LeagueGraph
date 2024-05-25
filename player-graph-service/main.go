package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(0) // This disables timestamp logging

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you've reached the Player Graph Service")
	// })

	// Fetch matches as part of the service start-up for testing
	puuid := os.Getenv("SOURCE_PUUID")
	region := os.Getenv("REGION")
	apiKey := os.Getenv("RIOT_API_KEY")

	matches := fetchMatches(puuid, region, apiKey)
	log.Println("Fetched matches: ", matches)

	uri := os.Getenv("NEO4J_CONN")
	username := os.Getenv("NEO4J_USER")
	password := os.Getenv("NEO4J_PASS")

	for _, match := range matches {
		matchData := fetchMatchData(match, region, apiKey)
		insertPlayerData(matchData, uri, username, password)
	}

	// wipeDatabase(uri, username, password)
}
