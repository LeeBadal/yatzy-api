package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"yatzy/internal/converters"
	"yatzy/internal/dbservice"
	"yatzy/internal/logic"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	router := gin.Default()
	router.Use(corsMiddleware())

	// Basic Authentication
	/*
		authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
			"admin": "admin123", // username : admin, password : admin123
		}))
	*/

	//authorizde but without authenication
	authorized := router.Group("/")

	authorized.POST("/create-game", corsMiddleware(), createGame)
	authorized.POST("/next-turn", corsMiddleware(), nextTurn)
	authorized.POST("submit-choice", corsMiddleware(), submitChoice)

	// Parse Static files
	//router.StaticFile("/", "./public/index.html")

	router.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

var targetAddress string

func init() {
	env := os.Getenv("ENV")
	if env == "development" {
		// Connect to localhost for development
		targetAddress = "localhost:50051"
	} else {
		// Connect to the remote server
		targetAddress = "dbservice-service.default.svc.cluster.local:50051"
	}
}

func createGame(c *gin.Context) {
	var req struct {
		NumPlayers int `json:"numPlayers" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(targetAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client using the connection.
	client := dbservice.NewDatabaseServiceClient(conn)

	// create game
	game := logic.CreateGame(req.NumPlayers)

	dbgame := converters.ConvertLogicGameStateToDBGameState(&game)

	// Call the AddGame RPC
	addGameRequest := dbservice.AddGameRequest{
		Uuid:      game.Uuid,
		GameState: dbgame,
	}

	// Copy the fields from addGameRequest to a new instance of AddGameRequest
	newAddGameRequest := &dbservice.AddGameRequest{
		Uuid:      addGameRequest.Uuid,
		GameState: addGameRequest.GameState,
	}

	addGameResponse, err := client.AddGame(context.Background(), newAddGameRequest)
	if err != nil {
		log.Fatalf("AddGame failed: %v", err)
	}
	log.Printf("AddGame Response: %v", addGameResponse)

	c.JSON(http.StatusOK, gin.H{"game": game})
}

// Takes uuid and index of dice to keep,
// calls dbservice, get game, then parses logic, then calls dbservice again adding the game once logic has been applied
// if no errors occur, returns updated game as response
func nextTurn(c *gin.Context) {
	var req struct {
		Keep []int  `json:"keep" binding:"required"`
		Uuid string `json:"uuid" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn, err := grpc.Dial(targetAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client using the connection.
	client := dbservice.NewDatabaseServiceClient(conn)

	GetGameRequest := &dbservice.GetGameRequest{
		Uuid: req.Uuid,
	}
	NewGetGameRequest := &dbservice.GetGameRequest{
		Uuid: GetGameRequest.Uuid,
	}

	getGameResponse, err := client.GetGame(context.Background(), NewGetGameRequest)
	if err != nil {
		log.Fatalf("GetGame failed: %v", err)
	}
	log.Printf("GetGame Response: %v", getGameResponse)

	gameval := getGameResponse.GetGameState()
	convertedGameState := converters.ConvertDBGameStateToLogicGameState(gameval)
	nextTurn, err := logic.NextTurn(*convertedGameState, req.Keep)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logic.PrintGameState(nextTurn)
	//convert nextturn back to dbgamestate
	dbgame := converters.ConvertLogicGameStateToDBGameState(&nextTurn)
	newAddGameRequest := &dbservice.AddGameRequest{
		Uuid:      getGameResponse.GetUuid(),
		GameState: dbgame,
	}

	addGameResponse, err := client.AddGame(context.Background(), newAddGameRequest)
	if err != nil {
		log.Fatalf("AddGame failed: %v", err)
	}
	log.Printf("AddGame Response: %v", addGameResponse)

	c.JSON(http.StatusOK, gin.H{"game": nextTurn})
}

func submitChoice(c *gin.Context) {
	var req struct {
		Choice string `json:"choice" binding:"required"`
		Uuid   string `json:"uuid" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate that choice is valid

	conn, err := grpc.Dial(targetAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client using the connection.
	client := dbservice.NewDatabaseServiceClient(conn)

	GetGameRequest := &dbservice.GetGameRequest{
		Uuid: req.Uuid,
	}
	NewGetGameRequest := &dbservice.GetGameRequest{
		Uuid: GetGameRequest.Uuid,
	}

	getGameResponse, err := client.GetGame(context.Background(), NewGetGameRequest)
	if err != nil {
		log.Fatalf("GetGame failed: %v", err)
	}
	log.Printf("GetGame Response: %v", getGameResponse)

	gameval := getGameResponse.GetGameState()
	convertedGameState := converters.ConvertDBGameStateToLogicGameState(gameval)
	convertedGameState.CategoryChoice = req.Choice
	nextTurn, err := logic.NextTurn(*convertedGameState, []int{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logic.PrintGameState(nextTurn)
	//convert nextturn back to dbgamestate
	dbgame := converters.ConvertLogicGameStateToDBGameState(&nextTurn)
	newAddGameRequest := &dbservice.AddGameRequest{
		Uuid:      getGameResponse.GetUuid(),
		GameState: dbgame,
	}

	addGameResponse, err := client.AddGame(context.Background(), newAddGameRequest)
	if err != nil {
		log.Fatalf("AddGame failed: %v", err)
	}
	log.Printf("AddGame Response: %v", addGameResponse)
	c.JSON(http.StatusOK, gin.H{"game": dbgame})
}

func CreateGamePtr(players int) *logic.GameState {
	game := logic.CreateGame(players)
	return &game
}
