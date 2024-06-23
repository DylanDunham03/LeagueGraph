// database.go

package main

import (
	"strconv"

	pb "github.com/DylanDunham03/LeagueGraph/player-graph-service/protos"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	"log"
	"time"
)

type Neo4jClient struct {
	Driver neo4j.Driver
}

func mapNeo4jNodeToPlayerProto(node neo4j.Node) pb.Player {
	LastSeen := ""
	if lastSeen, ok := node.Props["lastSeen"].(time.Time); ok {
		LastSeen = lastSeen.Format(time.RFC3339) // Converts time.Time to a RFC 3339 formatted string
	}

	return pb.Player{
		Puuid:          node.Props["puuid"].(string),
		RiotIdName:     node.Props["riotIdTagline"].(string),
		RiotIdGameName: node.Props["riotIdGameName"].(string),
		LastSeen:       LastSeen,
		Role:           node.Props["role"].(string),
	}
}

func mapNeo4jRelationshipToConnProto(r neo4j.Relationship, region string) pb.Connection {
	gameId := ""
	if id, ok := r.Props["gameId"].(int64); ok {
		gameId = strconv.FormatInt(id, 10)
	}

	timesPlayed := int32(-1)
	if tp, ok := r.Props["timesPlayed"].(int64); ok {
		timesPlayed = int32(tp)
	}

	LastPlayed := ""
	if lastPlayed, ok := r.Props["lastPlayed"].(time.Time); ok {
		LastPlayed = lastPlayed.Format(time.RFC3339)
	}

	return pb.Connection{
		GameId:      gameId,
		TimesPlayed: timesPlayed,
		LastPlayed:  LastPlayed,
		Region:      region,
	}
}

func NewNeo4jClient(uri, username, password string) (*Neo4jClient, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	return &Neo4jClient{Driver: driver}, nil
}

func (client *Neo4jClient) GetPlayerGraph(region string) (*pb.GraphResponse, error) {
	session := client.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	// Query for unique players
	playerQuery := `
        MATCH (p:Player)-[r:PLAYED_WITH {region: $region}]-()
        RETURN DISTINCT p
    `
	playerResult, err := session.Run(playerQuery, map[string]interface{}{"region": region})
	if err != nil {
		log.Printf("Error executing player query: %v", err)
		return nil, err
	}

	var response pb.GraphResponse
	playerMap := make(map[string]*pb.Player) // Map to store unique players

	for playerResult.Next() {
		record := playerResult.Record()
		node := record.Values[0].(neo4j.Node)
		player := mapNeo4jNodeToPlayerProto(node)
		response.Players = append(response.Players, &player)
		playerMap[player.Puuid] = &player // Add player to map for reference in connection processing
	}

	if err := playerResult.Err(); err != nil {
		log.Printf("Error processing player results: %v", err)
		return nil, err
	}

	// Query for unique connections
	connectionQuery := `
        MATCH (p:Player)-[r:PLAYED_WITH {region: $region}]-(q:Player)
        WHERE id(p) < id(q)
        RETURN p, r, q
    `
	connectionResult, err := session.Run(connectionQuery, map[string]interface{}{"region": region})
	if err != nil {
		log.Printf("Error executing connection query: %v", err)
		return nil, err
	}

	for connectionResult.Next() {
		record := connectionResult.Record()
		p := record.Values[0].(neo4j.Node)
		r := record.Values[1].(neo4j.Relationship)
		q := record.Values[2].(neo4j.Node)

		playerOne := mapNeo4jNodeToPlayerProto(p)
		playerTwo := mapNeo4jNodeToPlayerProto(q)
		connection := mapNeo4jRelationshipToConnProto(r, region)
		connection.PlayerOneUuid = playerOne.Puuid
		connection.PlayerTwoUuid = playerTwo.Puuid

		response.Connections = append(response.Connections, &connection)
	}

	if err := connectionResult.Err(); err != nil {
		log.Printf("Error processing connection results: %v", err)
		return nil, err
	}

	log.Printf("Total players processed: %d, Total connections processed: %d", len(response.Players), len(response.Connections))
	return &response, nil
}
