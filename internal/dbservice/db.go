package dbservice

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Game struct {
	ID        string                 `json:"id"`
	Version   int                    `json:"version"`
	GameState map[string]interface{} `json:"game_state"`
	CreatedAt string                 `json:"created_at"`
}

type Config struct {
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
}

// InitializeDB initializes the database connection.
func InitializeDB() (*sql.DB, error) {
	var config Config

	// Load the configuration file based on the environment
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
		configFile := fmt.Sprintf("config.%s.json", env)
		file, err := os.Open(configFile)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Decode the configuration file
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			return nil, err
		}
	} else if env == "production" {
		config.DBHost = os.Getenv("DB_HOST")
		config.DBPort = os.Getenv("DB_PORT")
		config.DBUser = os.Getenv("DB_USER")
		config.DBPassword = os.Getenv("DB_PASSWORD")
		config.DBName = os.Getenv("DB_NAME")
	} else {
		return nil, errors.New("invalid environment")

	}

	// Set the environment variables
	os.Setenv("DB_HOST", config.DBHost)
	os.Setenv("DB_PORT", config.DBPort)
	os.Setenv("DB_USER", config.DBUser)
	os.Setenv("DB_PASSWORD", config.DBPassword)
	os.Setenv("DB_NAME", config.DBName)

	// Create the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// AddGame adds a new game to the "games" table.
func AddGame(db *sql.DB, uuid string, gameState map[string]interface{}) error {
	// Convert the game state to JSON

	log.Printf("gameState: %v", gameState)
	gameStateJSON, err := json.Marshal(gameState)
	if err != nil {
		return err
	}

	log.Printf("gameStateJson: %v", gameStateJSON)
	// Prepare the SQL statement to insert a new game
	stmt, err := db.Prepare(`
        INSERT INTO games (id, game_state)
        VALUES ($1, $2)
    `)
	if err != nil {
		return err
	}

	// Execute the SQL statement
	_, err = stmt.Exec(uuid, gameStateJSON)
	if err != nil {
		return err
	}

	return nil
}

// GetGameByUUIDAndLatestVersion retrieves a game by its UUID and the latest version.
func GetGameByUUIDAndLatestVersion(db *sql.DB, uuid string) (*Game, error) {
	var game Game

	// Prepare the SQL statement to retrieve a game by UUID and latest version
	rows, err := db.Query(`
        SELECT id, version, game_state, created_at
        FROM games
        WHERE id = $1
        ORDER BY version DESC
        LIMIT 1
    `, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var gameStateJSON []byte

		err := rows.Scan(&game.ID, &game.Version, &gameStateJSON, &game.CreatedAt)
		if err != nil {
			return nil, err
		}

		// Unmarshal the JSON game state
		err = json.Unmarshal(gameStateJSON, &game.GameState)
		if err != nil {
			return nil, err
		}

		return &game, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// If no matching game is found, return an error or nil, depending on your use case
	return nil, sql.ErrNoRows
}

// function to connect to db and
