#!/bin/bash

# Navigate to the project root directory from scripts directory
cd "$(dirname "$0")"/..

echo "Building Go services..."
go build -o auth-service/auth-service ./auth-service
go build -o grpc-gateway/grpc-gateway ./grpc-gateway
go build -o player-graph-service/player-graph-service ./player-graph-service

echo "Starting services..."
# Start each service in the background and save their PIDs in the scripts folder
./auth-service/auth-service & echo $! > scripts/auth-service.pid
./grpc-gateway/grpc-gateway & echo $! > scripts/grpc-gateway.pid
./player-graph-service/player-graph-service & echo $! > scripts/player-graph-service.pid

echo "All services started."
