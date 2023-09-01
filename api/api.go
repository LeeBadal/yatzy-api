package api

import (
	"net/http"

	"yatzy/logic"

	"sync"

	"github.com/gin-gonic/gin"
)

type Event struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message chan string

	// New client connections
	NewClients chan chan string

	// Closed client connections
	ClosedClients chan chan string

	// Total client connections
	TotalClients map[chan string]bool
}

// New event messages are broadcast to all registered client connection channels
type ClientChan chan string

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
	router.StaticFile("/", "./public/index.html")

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

var (
	gameMutex sync.Mutex
	game      *logic.GameState
)

func createGame(c *gin.Context) {
	var req struct {
		NumPlayers int `json:"numPlayers" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game = CreateGamePtr(req.NumPlayers)
	logic.PrintGameState(*game)

	c.JSON(http.StatusOK, gin.H{"game": game})
}

func nextTurn(c *gin.Context) {
	var req struct {
		Keep []int `json:"keep" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gameMutex.Lock()
	defer gameMutex.Unlock()

	if game == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "game not found"})
		return
	}

	gameval := *game
	newGame, err := logic.NextTurn(gameval, req.Keep)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	*game = newGame
	logic.PrintGameState(*game)

	c.JSON(http.StatusOK, gin.H{"game": game})
}

func submitChoice(c *gin.Context) {
	var req struct {
		Choice string `json:"choice" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gameMutex.Lock()
	defer gameMutex.Unlock()
	if game == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "game not found"})
		return
	}
	game.CategoryChoice = req.Choice
	gameval := *game
	newGame, err := logic.NextTurn(gameval, []int{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	*game = newGame
	logic.PrintGameState(*game)
	c.JSON(http.StatusOK, gin.H{"game": game})
}

func CreateGamePtr(players int) *logic.GameState {
	game := logic.CreateGame(players)
	return &game
}
