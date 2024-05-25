package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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

type ParticipantDto struct {
	ChampionId         int    `json:"championId"`
	ChampionName       string `json:"championName"`
	IndividualPosition string `json:"individualPosition"`
	ParticipantId      int    `json:"participantId"`
	Puuid              string `json:"puuid"`
	RiotIdGameName     string `json:"riotIdGameName"`
	RiotIdTagline      string `json:"riotIdTagline"`
	Role               string `json:"role"`
	TeamId             int    `json:"teamId"`
	TeamPosition       string `json:"teamPosition"`
}

type MatchDto struct {
	Metadata struct {
		Participants []string `json:"participants"`
	} `json:"metadata"`
	InfoDto struct {
		GameDuration     int64            `json:"gameDuration"`
		GameEndTimestamp int64            `json:"gameEndTimestamp"`
		GameId           int64            `json:"gameId"`
		MapId            int64            `json:"mapId"`
		Participants     []ParticipantDto `json:"participants"`
	} `json:"info"`
}

func fetchMatchData(matchId string, region string, apiKey string) MatchDto {
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

	return match
}
func insertPlayerData(match MatchDto, uri string, username string, password string) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatal("Error creating Neo4j driver: ", err)
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	lastSeenTime := time.Unix(0, match.InfoDto.GameEndTimestamp*int64(time.Millisecond)).Format(time.RFC3339)

	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		// Check if the gameId already exists in any PLAYED_WITH relationship
		result, err := transaction.Run(
			`MATCH ()-[r:PLAYED_WITH {gameId: $gameId}]->() RETURN r`,
			map[string]interface{}{"gameId": match.InfoDto.GameId})
		if err != nil {
			log.Printf("Error checking for existing gameId: %v\n", err)
			return nil, err
		}

		// If the gameId is found, skip creating or updating nodes and relationships
		if result.Next() {
			log.Println("gameId already exists in PLAYED_WITH relationship, skipping update.")
			return nil, nil
		}

		// If the gameId does not exist, proceed with creating/updating nodes and relationships
		for _, participant := range match.InfoDto.Participants {
			_, err := transaction.Run(
				`MERGE (p:Player {puuid: $puuid})
				ON CREATE SET p.riotIdGameName = $riotIdGameName, p.riotIdTagline = $riotIdTagline, p.role = $role, p.lastSeen = datetime($lastSeenTime)
				ON MATCH SET p.riotIdGameName = $riotIdGameName, p.riotIdTagline = $riotIdTagline, p.role = $role, p.lastSeen = datetime($lastSeenTime)`,
				map[string]interface{}{
					"puuid":          participant.Puuid,
					"riotIdGameName": participant.RiotIdGameName,
					"riotIdTagline":  participant.RiotIdTagline,
					"role":           participant.Role,
					"lastSeenTime":   lastSeenTime,
				})
			if err != nil {
				log.Printf("Error creating/updating player: %v\n", err)
				return nil, err
			}
		}

		for i := 0; i < len(match.InfoDto.Participants); i++ {
			for j := i + 1; j < len(match.InfoDto.Participants); j++ {
				_, err = transaction.Run(
					`MATCH (p1:Player {puuid: $puuid1}), (p2:Player {puuid: $puuid2})
                     MERGE (p1)-[r:PLAYED_WITH]->(p2)
                     ON CREATE SET r.gameId = $gameId, r.timesPlayed = 1, r.lastPlayed = datetime($lastSeenTime)
                     ON MATCH SET r.timesPlayed = r.timesPlayed + 1, r.lastPlayed = datetime($lastSeenTime)`,
					map[string]interface{}{
						"puuid1":       match.InfoDto.Participants[i].Puuid,
						"puuid2":       match.InfoDto.Participants[j].Puuid,
						"gameId":       match.InfoDto.GameId,
						"lastSeenTime": lastSeenTime,
					})
				if err != nil {
					log.Printf("Error creating/updating relationship between %s and %s: %v\n", match.InfoDto.Participants[i].Puuid, match.InfoDto.Participants[j].Puuid, err)
					return nil, err
				}
			}
		}
		return nil, nil
	})
	if err != nil {
		log.Fatal("Error writing to Neo4j: ", err)
	}
}

func wipeDatabase(uri, username, password string) error {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return fmt.Errorf("error creating driver: %v", err) // start error messages with a lower case
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err = session.Run("MATCH (n) DETACH DELETE n", nil)
	if err != nil {
		return fmt.Errorf("error deleting all nodes and relationships: %v", err) // start error messages with a lower case
	}

	return nil
}
