package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/DylanDunham03/LeagueGraph/player-graph-service/protos" // import the protobuf package
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// func main() {
// 	log.SetFlags(0) // This disables timestamp logging

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	// Fetch matches as part of the service start-up for testing
// 	puuid := os.Getenv("SOURCE_PUUID")
// 	region := os.Getenv("REGION")
// 	apiKey := os.Getenv("RIOT_API_KEY")

// 	matches := fetchMatches(puuid, region, apiKey)
// 	log.Println("Fetched matches: ", matches)

// 	// uri := os.Getenv("NEO4J_CONN")
// 	// username := os.Getenv("NEO4J_USER")
// 	// password := os.Getenv("NEO4J_PASS")

// 	// for _, match := range matches {
// 	// 	matchData := fetchMatchData(match, region, apiKey)
// 	// 	insertPlayerData(matchData, region, uri, username, password)
// 	// }

// 	// wipeDatabase(uri, username, password)
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Neo4j client
	uri := os.Getenv("NEO4J_CONN")
	username := os.Getenv("NEO4J_USER")
	password := os.Getenv("NEO4J_PASS")
	neo4jClient, err := NewNeo4jClient(uri, username, password)
	if err != nil {
		log.Fatalf("Failed to create Neo4j client: %v", err)
	}

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50051") // Ensure this port matches across your services
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPlayerGraphServiceServer(s, &serverImpl{neo4jClient: neo4jClient})
	reflection.Register(s)
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type serverImpl struct {
	pb.UnimplementedPlayerGraphServiceServer
	neo4jClient *Neo4jClient // Use your Neo4jClient type here
}

// Implement the GetPlayerGraph gRPC method
func (s *serverImpl) GetPlayerGraph(ctx context.Context, req *pb.GraphRequest) (*pb.GraphResponse, error) {
	return s.neo4jClient.GetPlayerGraph(req.Region)
}
