package dbservice

import (
	"database/sql"
	"embed"

	"net"

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
	pb.RegisterDatabaseServiceServer(grpcServer, &DatabaseServiceServer{db: db})

	// Start the gRPC server
	lis, err := net.Listen("tcp", ":50051") // Use your desired port
	if err != nil {
		// Handle the error
	}
	if err := grpcServer.Serve(lis); err != nil {
		// Handle the error
	}

}
