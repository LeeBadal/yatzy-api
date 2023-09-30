// converters/converters.go

package converters

import (
	"encoding/json"
	"fmt"
	"log"
	db "yatzy/internal/dbservice"
	"yatzy/internal/logic"
)

// ConvertLogicGameStateToDBGameState converts a logic.GameState to a db.GameState.
func ConvertLogicGameStateToDBGameState(input *logic.GameState) *db.GameState {
	//convert logic gamestate to json
	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("JSON marshaling error:", err)
	}

	var dbGameState db.GameState
	if err := json.Unmarshal([]byte(jsonData), &dbGameState); err != nil {
		log.Fatalf("JSON unmarshal error: %v", err)
	}
	return &dbGameState
}

// ConvertDBGameStateToLogicGameState converts a db.GameState to a logic.GameState.
func ConvertDBGameStateToLogicGameState(input *db.GameState) *logic.GameState {
	//convert db gamestate to json
	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("JSON marshaling error:", err)
	}

	var logicGameState logic.GameState
	if err := json.Unmarshal([]byte(jsonData), &logicGameState); err != nil {
		log.Fatalf("JSON unmarshal error: %v", err)
	}
	return &logicGameState
}
