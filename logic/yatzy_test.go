package logic

import (
	"testing"
)

// test issubset
func TestIsSubset(t *testing.T) {
	dice := []int{1, 1, 2, 2, 3}
	subset := []int{1, 2}
	got := isSubset(dice, subset)
	want := true
	if got != want {
		t.Errorf("isSubset(%v, %v) = %v, want %v", dice, subset, got, want)
	}
}

// test issubset with empty subset
func TestIsSubset2(t *testing.T) {
	dice := []int{1, 1, 2, 2, 3}
	subset := []int{}
	got := isSubset(dice, subset)
	want := true
	if got != want {
		t.Errorf("isSubset(%v, %v) = %v, want %v", dice, subset, got, want)
	}
}

// test issubset with empty dice
func TestIsSubset3(t *testing.T) {
	dice := []int{}
	subset := []int{1, 2}
	got := isSubset(dice, subset)
	want := false
	if got != want {
		t.Errorf("isSubset(%v, %v) = %v, want %v", dice, subset, got, want)
	}
}

// test issubset with empty dice and subset
func TestIsSubset4(t *testing.T) {
	dice := []int{}
	subset := []int{}
	got := isSubset(dice, subset)
	want := true
	if got != want {
		t.Errorf("isSubset(%v, %v) = %v, want %v", dice, subset, got, want)
	}
}

// test issubset with empty dice and subset
func TestIsSubset5(t *testing.T) {
	dice := []int{1, 1, 2, 5, 3}
	subset := []int{5, 5, 5, 5, 5}
	got := isSubset(dice, subset)
	want := false
	if got != want {
		t.Errorf("isSubset(%v, %v) = %v, want %v", dice, subset, got, want)
	}
}

func TestNextTurn(t *testing.T) {
	game := CreateGame(2)
	game, _ = NextTurn(game, []int{0, 0, 0, 0, 0})
	got := game.CurrentPlayer
	want := 0
	if got != want {
		t.Errorf("nextTurn(%v) = %v, want %v", game, got, want)
	}
}

// Keep is not a subset so return current gamestate, rollsleft should be 2
func TestNextTurn2(t *testing.T) {
	game := CreateGame(2)
	game, err := NextTurn(game, []int{0, 0, 0, 0, 0})
	if err != nil {
		t.Errorf("nextTurn(%v) = err: %v", game, err)
	}
	got := game.RollsLeft
	want := 1
	if got != want {
		t.Errorf("nextTurn(%v) = %v, want %v", game, got, want)
	}
}

// Keep is not a subset so return current gamestate, rollsleft should be 2
func TestNextTurn3(t *testing.T) {
	game := CreateGame(2)
	want := game.Dice
	game, err := NextTurn(game, []int{1, 1, 1, 1, 1})
	if err != nil {
		t.Errorf("nextTurn(%v) = err: %v", game, err)
	}
	got := game.Dice
	if EqualSlice(got, want) != true {
		t.Errorf("nextTurn(%v) = %v, want %v", game, got, want)
	}
}

func EqualSlice(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
