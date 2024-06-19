// database.go

package main

import (
	"strconv"

	pb "github.com/DylanDunham03/LeagueGraph/player-graph-service/protos"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	// "log"
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

	result, err := session.Run(`
        MATCH (p:Player)-[r:PLAYED_WITH {region: $region}]-(q:Player)
        RETURN p, r, q
    `, map[string]interface{}{"region": region})
	if err != nil {
		return nil, err
	}

	var response pb.GraphResponse
	for result.Next() {
		record := result.Record()
		p := record.Values[0].(neo4j.Node)
		r := record.Values[1].(neo4j.Relationship)
		q := record.Values[2].(neo4j.Node)

		// Assuming a function mapNeo4jNodeToPlayerProto exists and converts Neo4j Node to Player protobuf
		playerOne := mapNeo4jNodeToPlayerProto(p)
		playerTwo := mapNeo4jNodeToPlayerProto(q)
		connection := mapNeo4jRelationshipToConnProto(r, region)
		connection.PlayerOneUuid = playerOne.Puuid
		connection.PlayerTwoUuid = playerTwo.Puuid

		response.Players = append(response.Players, &playerOne, &playerTwo)
		response.Connections = append(response.Connections, &connection)
	}
	return &response, nil
}
