package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	pb "github.com/DylanDunham03/LeagueGraph/player-graph-service/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterPlayerGraphServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		return err
	}

	// Setup CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // This will allow any domain, adjust if necessary
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Or you can specify headers you want to allow
		AllowCredentials: true,
		Debug:            true, // Shows detailed logs of CORS operations
	}).Handler(mux)

	return http.ListenAndServe(":8080", corsHandler)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err)
	}
}
