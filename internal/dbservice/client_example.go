package dbservice

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func test() {
	// Establish a gRPC connection to your server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := NewDatabaseServiceClient(conn)

	// Make a customized gRPC call
	gameUUID := "your_game_uuid"
	response, err := client.GetGame(context.Background(), &GetGameRequest{Uuid: gameUUID})
	if err != nil {
		log.Fatalf("Error calling GetGame: %v", err)
	}

	// Handle the response
	fmt.Printf("Game UUID: %s\n", response.GetUuid())
	// Process other fields in the response as needed
}
