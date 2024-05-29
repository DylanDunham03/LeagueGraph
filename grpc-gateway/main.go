package main

import (
	// "context"
	"log"
	"net"

	pb "github.com/DylanDunham03/LeagueGraph/player-graph-service/protos"

	"google.golang.org/grpc"
	// "github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type serverImpl struct {
	pb.UnimplementedPlayerGraphServiceServer
}

// func (s *serverImpl) GetPlayerGraph(ctx context.Context, req *pb.GraphRequest) (*pb.GraphResponse, error) {
// 	// Implementation of fetching data from Neo4j and returning it
// }

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPlayerGraphServiceServer(s, &serverImpl{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
