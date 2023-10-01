package dbservice

import (
	"database/sql"
	"embed"
	"encoding/json"
	"log"

	"net"

	"context"

	"github.com/pressly/goose/v3"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	port = ":50051"
)

var embedMigrations embed.FS

type dbServiceServer struct {
	UnimplementedDatabaseServiceServer
}

func Server() {
	var db *sql.DB
	// setup database

	//goose.SetBaseFS(embedMigrations)

	db, err := InitializeDB()
	if err != nil {
		panic(err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	//run migrations
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	// Start the gRPC server
	lis, err := net.Listen("tcp", port) // Use your desired port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register your DatabaseServiceServer
	RegisterDatabaseServiceServer(grpcServer, &dbServiceServer{})

	log.Printf("Starting gRPC listener on port " + port)
	log.Printf("gRPC server listening at %v", lis.Addr())

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}

}

func (s *dbServiceServer) AddGame(ctx context.Context, in *AddGameRequest) (*AddGameResponse, error) {
	// Convert the game state to JSON format

	gameStateJSON, err := protojson.Marshal(in.GetGameState())
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into a map[string]interface{}
	var gameState map[string]interface{}
	if err := json.Unmarshal(gameStateJSON, &gameState); err != nil {
		return nil, err
	}

	// open db connection
	db, err := InitializeDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Add the game to the database
	err = AddGame(db, in.GetUuid(), gameState)
	if err != nil {
		// Handle the error
		return nil, err
	}

	// Return a success response
	return &AddGameResponse{Success: true}, nil
}

func (s *dbServiceServer) GetGame(ctx context.Context, in *GetGameRequest) (*GetGameResponse, error) {
	// Get the game from the database
	// open db connection
	db, err := InitializeDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	game, err := GetGameByUUIDAndLatestVersion(db, in.GetUuid())

	log.Printf("game: %v", game)

	if err != nil {
		// Handle the error
		return nil, err
	}

	// Convert the game state to a byte slice
	gameStateBytes, err := json.Marshal(game.GameState)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into a new GameState struct
	var gameState GameState
	if err := json.Unmarshal(gameStateBytes, &gameState); err != nil {
		return nil, err
	}

	// Return the game
	return &GetGameResponse{
		Uuid:      game.ID,
		GameState: &gameState,
	}, nil
}
