package dbservice

import (
	"database/sql"
	"embed"

	"net"

	"context"

	"github.com/pressly/goose/v3"
	grpc "google.golang.org/grpc"
)

var embedMigrations embed.FS

func main() {
	var db *sql.DB
	// setup database

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register your DatabaseServiceServer
	RegisterDatabaseServiceServer(grpcServer, &UnimplementedDatabaseServiceServer{})

	// Start the gRPC server
	lis, err := net.Listen("tcp", ":50051") // Use your desired port
	if err != nil {
		// Handle the error
	}
	if err := grpcServer.Serve(lis); err != nil {
		// Handle the error
	}

}

type dbServiceServer struct {
	UnimplementedDatabaseServiceServer
}

// add game to db
func (s *dbServiceServer) AddGame(ctx context.Context, in *AddGameRequest, db *sql.DB) (*AddGameResponse, error) {

	// Convert the game state to map[string]interface{}
	gameState := make(map[string]interface{})
	for k, v := range in.GetGameState() {
		gameState[k] = v
	}

	// Add the game to the database
	err := AddGame(db, in.GetUuid(), gameState)
	for k, v := range in.GetGameState() {
		gameState[k] = v
	}

	if err != nil {
		// Handle the error
	}

	// Return a success response
	return &AddGameResponse{}, nil
}

// get game from db
func (s *dbServiceServer) GetGame(ctx context.Context, in *GetGameRequest, db *sql.DB) (*GetGameResponse, error) {

	// Get the game from the database
	game, err := GetGameByUUIDAndLatestVersion(db, in.GetUuid())
	if err != nil {
		// Handle the error
	}

	// Convert the game state to map[string]string
	gameState := make(map[string]string)
	for k, v := range game.GameState {
		gameState[k] = v.(string)
	}

	// Return the game
	return &GetGameResponse{
		Uuid:      game.ID,
		GameState: gameState,
	}, nil
}
