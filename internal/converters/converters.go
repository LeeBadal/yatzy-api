// converters/converters.go

package converters

import (
	"encoding/json"
	"fmt"
	"log"
	db "yatzy/internal/dbservice"
	"yatzy/internal/logic"
)

const (
//list containing all keys like "ones" "twos" etc"

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

	keyList := []string{"Ones", "Twos", "Threes", "Fours", "Fives", "Sixes", "OnePair", "TwoPairs", "ThreeOfAKind", "FourOfAKind", "SmallStraight", "LargeStraight", "FullHouse", "Chance", "Yatzy"}

	var logicGameState logic.GameState
	if err := json.Unmarshal([]byte(jsonData), &logicGameState); err != nil {
		log.Fatalf("JSON unmarshal error: %v", err)
	}

	for x := 0; x < len(logicGameState.Players); x++ {
		for _, key := range keyList {
			if _, exists := logicGameState.Players[x].Score[key]; exists {
				fmt.Printf("Key %s exists in the map\n", key)
			} else {
				fmt.Printf("Key %s does not exist in the map\n", key)
				// if it does not exist, add it to the map with value 0
				logicGameState.Players[x].Score[key] = 0
			}
		}
	}

	fmt.Printf("%v", logicGameState.Players[0].Score)

	return &logicGameState
}
