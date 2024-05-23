package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func fetchMatches(puuid string, region string, apiKey string) []string {
	if puuid == "" || region == "" || apiKey == "" {
		log.Fatal("PUUID, region, and API key must be provided")
	}

	url := "https://" + region + ".api.riotgames.com/lol/match/v5/matches/by-puuid/" + puuid + "/ids?api_key=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching matches: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	log.Println("Raw API Response:", string(body))

	var matches []string
	if err := json.Unmarshal(body, &matches); err != nil {
		log.Fatal("Error parsing JSON: ", err)
	}

	return matches
}

type MatchDto struct {
	Metadata struct {
		Participants []string `json:"participants"`
	} `json:"metadata"`
	// Other fields omitted for brevity
}

func fetchMatchData(matchId string, region string, apiKey string) []string {
	if matchId == "" || region == "" || apiKey == "" {
		log.Fatal("matchId, region, and API key must be provided")
	}
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/%s?api_key=%s", region, matchId, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch match data: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response body: ", err)
	}

	var match MatchDto
	if err := json.Unmarshal(body, &match); err != nil {
		log.Fatal("Failed to unmarshal JSON: ", err)
	}

	return match.Metadata.Participants
}
