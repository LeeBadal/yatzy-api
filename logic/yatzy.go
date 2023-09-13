package logic

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
	"yatzy/validators"

	"github.com/google/uuid"
)

type Player struct {
	Name  string
	Score map[string]int
}

type GameState struct {
	Players         []Player
	CurrentPlayer   int
	RollsLeft       int
	Dice            []int
	RoundsLeft      int
	CategoryChoice  string
	ScoreCalculator map[string]int
	Uuid            string
}

// takes an array of ints, either 0 or 1. If the index is 1 the dice is kept, if it is 0 it is rolled
func rollDice(diceState []int, old []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	newDice := make([]int, len(old))
	copy(newDice, old)

	for i := 0; i < len(old); i++ {
		if diceState[i] == 0 {
			newDice[i] = r.Intn(6) + 1
		}
	}
	return newDice
}

/*func rollDice(numberOfDice []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	roll := make([]int, numberOfDice)
	for i := 0; i < numberOfDice; i++ {
		roll[i] = r.Intn(6) + 1
		println(roll[i])
	}
	return roll
}*/

// new turn, return the copy of the current game state, current player should be the opposite of the current player
func nextPlayer(gameState GameState) GameState {
	if gameState.RoundsLeft == 0 {
		return gameState
	}

	if gameState.CurrentPlayer == len(gameState.Players)-1 {
		gameState.RoundsLeft--
	}
	gameState.CurrentPlayer = (gameState.CurrentPlayer + 1) % len(gameState.Players)
	gameState.CategoryChoice = ""
	gameState.RollsLeft = 2
	gameState.Dice = rollDice([]int{0, 0, 0, 0, 0}, gameState.Dice)
	gameState = whatIfScore(gameState)
	return gameState
}

// get a print dice
func printDice(game GameState) {
	dice := game.Dice
	var s string
	for i := 0; i < len(dice); i++ {
		s += strconv.Itoa(int(dice[i])) + " "
	}
	println(s)
}

// print score, players, current player, rolls left, dice in a nice format
func PrintGameState(game GameState) {
	println("Current player: " + game.Players[game.CurrentPlayer].Name)
	println("Rolls left: " + strconv.Itoa(game.RollsLeft))
	printDice(game)
}

// check if array a is subset of b
func isSubset(b []int, a []int) bool {
	bCount := make(map[int]int)
	for _, x := range b {
		bCount[x]++
	}
	for _, x := range a {
		if bCount[x] == 0 {
			return false
		}
		bCount[x]--
	}
	return true
}

// whatifScore helper
func whatIfScore(gameState GameState) GameState {
	dice := make([]int, len(gameState.Dice))
	copy(dice, gameState.Dice)
	gameState.ScoreCalculator = validators.CalculateScores(dice, gameState.Players[gameState.CurrentPlayer].Score)
	return gameState
}

// check if keep is a subset of gamestate, if it is, then roll the dice that are not in keep
// if rollsleft is 0, then update score, a user must choose a category, if not gamestate will be returned unchanged
func NextTurn(gameState GameState, keep []int) (GameState, error) {

	if gameState.RollsLeft == 0 || gameState.CategoryChoice != "" {
		gameState = whatIfScore(gameState)
		if gameState.CategoryChoice == "" {
			println("You must make a choice")
			return gameState, errors.New("You must make a choice")
		}
		gameState, err := validateChoice(gameState)
		if err != nil {
			return gameState, errors.New("Invalid choice")
		}
		return nextPlayer(gameState), nil
	}

	gameState.Dice = rollDice(keep, gameState.Dice)
	gameState.RollsLeft = gameState.RollsLeft - 1
	gameState = whatIfScore(gameState)
	return gameState, nil
}

// validate choice and update score
func validateChoice(gameState GameState) (GameState, error) {
	//checks if users choice is not already chosen
	if gameState.Players[gameState.CurrentPlayer].Score[gameState.CategoryChoice] != -1 {
		println("Category already chosen")
		return gameState, errors.New("Category already chosen or not valid")
	}
	gameState.Players[gameState.CurrentPlayer].Score[gameState.CategoryChoice] = gameState.ScoreCalculator[gameState.CategoryChoice]
	gameState.CategoryChoice = ""
	return gameState, nil
}

///INITIALIZERS

func CreateGame(players int) GameState {
	gameState := GameState{}
	gameState.Players = make([]Player, players)
	for i := 0; i < players; i++ {
		gameState.Players[i] = CreatePlayer("Player " + strconv.Itoa(i))
	}
	gameState.CurrentPlayer = 0
	gameState.RollsLeft = 2
	gameState.Dice = rollDice([]int{0, 0, 0, 0, 0}, make([]int, 5))
	gameState.ScoreCalculator = validators.CalculateScores(gameState.Dice, gameState.Players[gameState.CurrentPlayer].Score)
	gameState.CategoryChoice = ""
	gameState.RoundsLeft = len(gameState.Players[0].Score)
	gameState.Uuid = uuid.New().String()
	return gameState
}

func CreatePlayer(name string) Player {
	player := Player{}
	player.Name = name
	player.Score = map[string]int{
		"Ones":   -1,
		"Twos":   -1,
		"Threes": -1,
		"Fours":  -1,
		"Fives":  -1,
		"Sixes":  -1,

		"OnePair":       -1,
		"TwoPairs":      -1,
		"ThreeOfAKind":  -1,
		"FourOfAKind":   -1,
		"SmallStraight": -1,
		"LargeStraight": -1,
		"FullHouse":     -1,
		"Chance":        -1,
		"Yatzy":         -1,
	}

	return player
}
