#!/bin/bash

# Navigate to the project root directory from scripts directory
cd "$(dirname "$0")"/..

# Function to kill a service using its PID file
kill_service() {
    pid_file="scripts/$1"
    if [ -f "$pid_file" ]; then
        pid=$(cat "$pid_file")
        echo "Stopping service with PID $pid..."
        kill "$pid" && rm "$pid_file"
    else
        echo "PID file not found for $pid_file, service might not be running."
    fi
}

# Kill each service
kill_service "auth-service.pid"
kill_service "grpc-gateway.pid"
kill_service "player-graph-service.pid"

echo "All services stopped."
