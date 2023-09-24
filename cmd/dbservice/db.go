package dbservice

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Game struct {
	ID        string                 `json:"id"`
	Version   int                    `json:"version"`
	GameState map[string]interface{} `json:"game_state"`
	CreatedAt string                 `json:"created_at"`
}

// InitializeDB initializes the database connection.
func InitializeDB() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// AddGame adds a new game to the "games" table.
func AddGame(db *sql.DB, uuid string, gameState map[string]interface{}) error {
	// Convert the game state to JSON
	gameStateJSON, err := json.Marshal(gameState)
	if err != nil {
		return err
	}

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
