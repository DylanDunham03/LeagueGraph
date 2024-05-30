package main

import (
    "context"
    "log"
    "net"
    "os"

    pb "github.com/DylanDunham03/LeagueGraph/player-graph-service/protos"
    "google.golang.org/grpc"
    "github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type serverImpl struct {
    pb.UnimplementedPlayerGraphServiceServer
    neo4jDriver neo4j.Driver
}

func (s *serverImpl) GetPlayerGraph(ctx context.Context, req *pb.GraphRequest) (*pb.GraphResponse, error) {
    session := s.neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
    defer session.Close()

    result, err := session.Run(`
        MATCH (p:Player)-[r:PLAYED_WITH]-(q:Player)
        WHERE p.region = $region OR q.region = $region
        RETURN p, r, q
    `, map[string]interface{}{"region": req.Region})
    if err != nil {
        return nil, err
    }

    var response pb.GraphResponse
    for result.Next() {
        record := result.Record()
        p := record.Values[0].(neo4j.Node)
        r := record.Values[1].(neo4j.Relationship)
        q := record.Values[2].(neo4j.Node)

        playerOne := pb.Player{
            Puuid:         p.Props["puuid"].(string),
            RiotIdName:    p.Props["riotIdName"].(string),
            RiotIdGameName: p.Props["riotIdGameName"].(string),
            LastSeen:      p.Props["lastSeen"].(string),
            Role:          p.Props["role"].(string),
        }
        playerTwo := pb.Player{
            Puuid:         q.Props["puuid"].(string),
            RiotIdName:    q.Props["riotIdName"].(string),
            RiotIdGameName: q.Props["riotIdGameName"].(string),
            LastSeen:      q.Props["lastSeen"].(string),
            Role:          q.Props["role"].(string),
        }
        connection := pb.Connection{
            PlayerOneUuid: playerOne.Puuid,
            PlayerTwoUuid: playerTwo.Puuid,
            GameId:        r.Props["gameId"].(string),
            TimesPlayed:   int32(r.Props["timesPlayed"].(int)),
            LastPlayed:    r.Props["lastPlayed"].(string),
        }

        response.Players = append(response.Players, &playerOne, &playerTwo)
        response.Connections = append(response.Connections, &connection)
    }

    return &response, nil
}

func main() {
    uri := os.Getenv("NEO4J_CONN")
    username := os.Getenv("NEO4J_USER")
    password := os.Getenv("NEO4J_PASS")
    driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
    if err != nil {
        log.Fatalf("Error creating Neo4j driver: %v", err)
    }
    defer driver.Close()

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterPlayerGraphServiceServer(grpcServer, &serverImpl{neo4jDriver: driver})
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
