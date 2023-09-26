package validators

import "testing"

func TestIsPair(t *testing.T) {
	dice := []int{1, 1, 2, 2, 3}
	got := Pair(dice)
	want := 4
	if got != want {
		t.Errorf("isTwoPair(%v) = %v, want %v", dice, got, want)
	}
}

// write more unit tests for testisTwoPair here
func TestIsPair2(t *testing.T) {
	dice := []int{1, 1, 2, 2, 2}
	got := Pair(dice)
	want := 4
	if got != want {
		t.Errorf("isTwoPair(%v) = %v, want %v", dice, got, want)
	}
}

func TestIsPair3(t *testing.T) {
	dice := []int{1, 1, 2, 3, 4}
	got := Pair(dice)
	want := 2
	if got != want {
		t.Errorf("isTwoPair(%v) = %v, want %v", dice, got, want)
	}
}
func TestIsPair4(t *testing.T) {
	dice := []int{1, 5, 2, 3, 4}
	got := Pair(dice)
	want := 0
	if got != want {
		t.Errorf("isTwoPair(%v) = %v, want %v", dice, got, want)
	}
}

func TestThreeOfAKind(t *testing.T) {
	dice := []int{1, 1, 1, 2, 3}
	got := Toak(dice)
	want := 3
	if got != want {
		t.Errorf("isThreeOfAKind(%v) = %v, want %v", dice, got, want)
	}
}

func TestThreeOfAKind2(t *testing.T) {
	dice := []int{1, 1, 1, 1, 3}
	got := Toak(dice)
	want := 3
	if got != want {
		t.Errorf("isThreeOfAKind(%v) = %v, want %v", dice, got, want)
	}
}

func TestThreeOfAKind3(t *testing.T) {
	dice := []int{1, 1, 3, 2, 2}
	got := Toak(dice)
	want := 0
	if got != want {
		t.Errorf("isThreeOfAKind(%v) = %v, want %v", dice, got, want)
	}
}

func TestFourOfAKind(t *testing.T) {
	dice := []int{1, 1, 1, 1, 3}
	got := Foak(dice)
	want := 4
	if got != want {
		t.Errorf("isFourOfAKind(%v) = %v, want %v", dice, got, want)
	}
}

func TestFourOfAKind2(t *testing.T) {
	dice := []int{1, 1, 1, 1, 1}
	got := Foak(dice)
	want := 4
	if got != want {
		t.Errorf("isFourOfAKind(%v) = %v, want %v", dice, got, want)
	}
}

func TestFourOfAKind3(t *testing.T) {
	dice := []int{1, 1, 1, 2, 3}
	got := Foak(dice)
	want := 0
	if got != want {
		t.Errorf("isFourOfAKind(%v) = %v, want %v", dice, got, want)
	}
}

func TestTwoPairs(t *testing.T) {
	dice := []int{1, 1, 2, 2, 3}
	got := TwoPairs(dice)
	want := 6
	if got != want {
		t.Errorf("isTwoPairs(%v) = %v, want %v", dice, got, want)
	}
}

func TestTwoPairs2(t *testing.T) {
	dice := []int{1, 1, 2, 2, 2}
	got := TwoPairs(dice)
	want := 6
	if got != want {
		t.Errorf("isTwoPairs(%v) = %v, want %v", dice, got, want)
	}
}

func TestTwoPairs3(t *testing.T) {
	dice := []int{1, 1, 2, 3, 4}
	got := TwoPairs(dice)
	want := 0
	if got != want {
		t.Errorf("isTwoPairs(%v) = %v, want %v", dice, got, want)
	}
}

func TestTwoPairs4(t *testing.T) {
	dice := []int{2, 2, 2, 2, 3}
	got := TwoPairs(dice)
	want := 0
	if got != want {
		t.Errorf("isTwoPairs(%v) = %v, want %v", dice, got, want)
	}
}

func TestEqualSlice(t *testing.T) {
	dice := []int{1, 1, 2, 3, 4}
	got := EqualSlice(dice, dice)
	want := true
	if got != want {
		t.Errorf("isTwoPairs(%v) = %v, want %v", dice, got, want)
	}
}

func TestEqualSlice2(t *testing.T) {
	dice := []int{1, 1, 2, 3, 4}
	dice2 := []int{1, 1, 2, 3, 6}
	got := EqualSlice(dice, dice2)
	want := false
	if got != want {
		t.Errorf("isTwoPairs(%v) = %v, want %v", dice, got, want)
	}
}

func TestSmallStraight(t *testing.T) {
	dice := []int{1, 2, 3, 4, 5}
	got := SmallStraight(dice)
	want := 15
	if got != want {
		t.Errorf("isSmallStraight(%v) = %v, want %v", dice, got, want)
	}
}

func TestSmallStraight2(t *testing.T) {
	dice := []int{1, 2, 3, 4, 6}
	got := SmallStraight(dice)
	want := 0
	if got != want {
		t.Errorf("isSmallStraight(%v) = %v, want %v", dice, got, want)
	}
}

func TestLargeStraight(t *testing.T) {
	dice := []int{2, 3, 4, 5, 6}
	got := LargeStraight(dice)
	want := 20
	if got != want {
		t.Errorf("isLargeStraight(%v) = %v, want %v", dice, got, want)
	}
}

func TestLargeStraight2(t *testing.T) {
	dice := []int{1, 2, 3, 4, 6}
	got := LargeStraight(dice)
	want := 0
	if got != want {
		t.Errorf("isLargeStraight(%v) = %v, want %v", dice, got, want)
	}
}

func TestFullHouse(t *testing.T) {
	dice := []int{1, 1, 2, 2, 2}
	got := FullHouse(dice)
	want := 8
	if got != want {
		t.Errorf("isFullHouse(%v) = %v, want %v", dice, got, want)
	}
}

func TestFullHouse2(t *testing.T) {
	dice := []int{1, 1, 2, 2, 3}
	got := FullHouse(dice)
	want := 0
	if got != want {
		t.Errorf("isFullHouse(%v) = %v, want %v", dice, got, want)
	}
}

func TestFullHouse3(t *testing.T) {
	dice := []int{1, 1, 1, 2, 2}
	got := FullHouse(dice)
	want := 7
	if got != want {
		t.Errorf("isFullHouse(%v) = %v, want %v", dice, got, want)
	}
}

func TestFullHouse4(t *testing.T) {
	dice := []int{1, 1, 1, 1, 1}
	got := FullHouse(dice)
	want := 0
	if got != want {
		t.Errorf("isFullHouse(%v) = %v, want %v", dice, got, want)
	}
}

func TestYatzy(t *testing.T) {
	dice := []int{1, 1, 1, 1, 1}
	got := Yatzy(dice)
	want := 55
	if got != want {
		t.Errorf("isYatzy(%v) = %v, want %v", dice, got, want)
	}
}

func TestYatzy2(t *testing.T) {
	dice := []int{1, 1, 1, 1, 2}
	got := Yatzy(dice)
	want := 0
	if got != want {
		t.Errorf("isYatzy(%v) = %v, want %v", dice, got, want)
	}
}

func TestSum(t *testing.T) {
	dice := []int{1, 1, 1, 1, 2}
	got := sum(dice)
	want := 6
	if got != want {
		t.Errorf("isSum(%v) = %v, want %v", dice, got, want)
	}
}

func TestSum2(t *testing.T) {
	dice := []int{1, 1, 1, 1, 1}
	got := sum(dice)
	want := 5
	if got != want {
		t.Errorf("isSum(%v) = %v, want %v", dice, got, want)
	}
}

func TestOnes(t *testing.T) {
	dice := []int{1, 1, 1, 1, 2}
	got := Ones(dice)
	want := 4
	if got != want {
		t.Errorf("isOnes(%v) = %v, want %v", dice, got, want)
	}
}

func TestOnes2(t *testing.T) {
	dice := []int{2, 2, 2, 2, 2}
	got := Ones(dice)
	want := 0
	if got != want {
		t.Errorf("isOnes(%v) = %v, want %v", dice, got, want)
	}
}

func TestBonus(t *testing.T) {
	dice := map[string]int{
		"Bonus":         -1,
		"Chance":        -1,
		"Fives":         -1,
		"FourOfAKind":   -1,
		"Fours":         -1,
		"FullHouse":     -1,
		"LargeStraight": -1,
		"OnePair":       -1,
		"Ones":          -1,
		"Sixes":         -1,
		"SmallStraight": -1,
		"ThreeOfAKind":  -1,
		"Threes":        -1,
		"TwoPairs":      -1,
		"Twos":          -1,
		"Yatzy":         -1,
	}
	got := Bonus(dice)
	want := 0
	if got != want {
		t.Errorf("isBonus(%v) = %v, want %v", dice, got, want)
	}
}

func TestBonus2(t *testing.T) {
	dice := map[string]int{
		"Chance":        -1,
		"Fives":         15,
		"FourOfAKind":   -1,
		"Fours":         12,
		"FullHouse":     -1,
		"LargeStraight": -1,
		"OnePair":       -1,
		"Ones":          3,
		"Sixes":         18,
		"SmallStraight": -1,
		"ThreeOfAKind":  -1,
		"Threes":        9,
		"TwoPairs":      -1,
		"Twos":          6,
		"Yatzy":         -1,
	}
	got := Bonus(dice)
	want := 50
	if got != want {
		t.Errorf("isBonus(%v) = %v, want %v", dice, got, want)
	}
}

func TestTotal(t *testing.T) {
	dice := map[string]int{
		"Chance":        -1,
		"Fives":         15,
		"FourOfAKind":   -1,
		"Fours":         12,
		"FullHouse":     -1,
		"LargeStraight": -1,
		"OnePair":       -1,
		"Ones":          3,
		"Sixes":         18,
		"SmallStraight": -1,
		"ThreeOfAKind":  -1,
		"Threes":        9,
		"TwoPairs":      -1,
		"Twos":          6,
		"Yatzy":         -1,
	}
	got := Total(dice)
	want := 63
	if got != want {
		t.Errorf("isTotal(%v) = %v, want %v", dice, got, want)
	}
}
