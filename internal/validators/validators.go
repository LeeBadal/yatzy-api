// a package of validators, the validators take an int array and return a bool if the array is valid against that score
// eg. the validator for 3 of a kind will return true if the array contains 3 of the same number
package validators

import (
	"sort"
)

// check if array contains at least a two pair, return sum if pair exists otherwise 0
// iterate backwards through sorted array, if two numbers are equal, return sum of pair
func Pair(dice []int) int {
	//make copy of dice
	sort.Ints(dice)
	pairSum := 0
	for i := len(dice) - 1; i > 0; i-- {
		if dice[i] == dice[i-1] {
			pairSum = dice[i] * 2
			break
		}
	}
	return pairSum
}

// TOAK is abbreviation for Three of a kind
// check if array contains at least a three pair, return sum if three of a kind exists
// iterate backwards through sorted array, if three numbers are equal, return sum of three of a kind
func Toak(dice []int) int {
	sort.Ints(dice)
	sum := 0
	for i := len(dice) - 1; i > 1; i-- {
		if dice[i] == dice[i-1] && dice[i] == dice[i-2] {
			sum = dice[i] * 3
			break
		}
	}
	return sum
}

// FOAK is abbreviation for Four of a kind
// check if array contains at least a four pair, return sum if four of a kind exists
// iterate backwards through sorted array, if four numbers are equal, return sum of four of a kind
func Foak(dice []int) int {
	sort.Ints(dice)
	sum := 0
	for i := len(dice) - 1; i > 2; i-- {
		if dice[i] == dice[i-1] && dice[i] == dice[i-2] && dice[i] == dice[i-3] {
			sum = dice[i] * 4
			break
		}
	}
	return sum
}

// check if there are two pairs, return sum of two highest pairs
// iterate backwkards through sorted array, if two numbers are equal, add to sum
func TwoPairs(dice []int) int {
	sort.Ints(dice)
	sum := 0
	pairs := 0
	pairNumber := 0
	for i := len(dice) - 1; i > 0; i-- {
		if dice[i] == dice[i-1] && dice[i] != pairNumber {
			sum += dice[i] * 2
			pairs++
			i--
			pairNumber = dice[i]
		}
		if pairs == 2 {
			return sum
		}
	}
	return 0
}

// Checks if slice are equal, O(n)... because
func EqualSlice(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	emap := make(map[int]int)
	for i := 0; i < len(a); i++ {
		emap[a[i]]++
	}

	for _, x := range b {
		if c, f := emap[x]; f {
			if c == 1 {
				delete(emap, x)
			} else {
				emap[x]--
			}
		} else {
			return false
		}

	}
	return true
}

// check small straight using EqualSlice
func SmallStraight(dice []int) int {
	sort.Ints(dice)
	if EqualSlice(dice, []int{1, 2, 3, 4, 5}) {
		return 15
	}
	return 0
}

// check large straight using EqualSlice
func LargeStraight(dice []int) int {
	sort.Ints(dice)
	if EqualSlice(dice, []int{2, 3, 4, 5, 6}) {
		return 20
	}
	return 0
}

func remove(s []int, i int) []int {
	for j := 0; j < len(s); j++ {
		if s[j] == i {
			copy(s[j:], s[j+1:])
			s = s[:len(s)-1]
			break
		}
	}
	return s
}

// check full house, full house, one pair and one three of a kind
func FullHouse(dice []int) int {
	toak := Toak(dice)
	if toak > 0 {
		for i := 0; i < 2; i++ {
			dice = remove(dice, toak/3)
		}
		pair := Pair(dice)
		if pair > 0 {
			if pair/2 != toak/3 {
				return pair + toak
			}
		}
	}
	return 0
}

// 6 3 6 6 6
// 12
//

func sum(dice []int) int {
	sum := 0
	for i := 0; i < len(dice); i++ {
		sum += dice[i]
	}
	return sum
}

func Chance(dice []int) int {
	return sum(dice)
}

// check yatzy using EqualSlice
func Yatzy(dice []int) int {
	ssum := sum(dice)
	if dice[0]*len(dice) == ssum {
		return 50 + sum(dice)
	}
	return 0
}

func Upper(dice []int, number int) int {
	sum := 0
	for i := 0; i < len(dice); i++ {
		if dice[i] == number {
			sum += number
		}
	}
	return sum
}

func Ones(dice []int) int {
	return Upper(dice, 1)
}
func Twos(dice []int) int {
	return Upper(dice, 2)
}
func Threes(dice []int) int {
	return Upper(dice, 3)
}
func Fours(dice []int) int {
	return Upper(dice, 4)
}
func Fives(dice []int) int {
	return Upper(dice, 5)
}
func Sixes(dice []int) int {
	return Upper(dice, 6)
}

func Bonus(currentPlayerScore map[string]int) int {
	sum := 0
	upperKeys := []string{"Ones", "Twos", "Threes", "Fours", "Fives", "Sixes"}
	for _, key := range upperKeys {
		sum += currentPlayerScore[key]
	}
	if sum >= 63 {
		return 50
	}
	return 0
}

func Total(currentPlayerScore map[string]int) int {
	sum := 0
	for _, value := range currentPlayerScore {
		if value != -1 {
			sum += value
		}
	}
	return sum
}

func CalculateScores(dice []int, currentPlayerScore map[string]int) map[string]int {
	scores := make(map[string]int)
	scores["Ones"] = Ones(dice)
	scores["Twos"] = Twos(dice)
	scores["Threes"] = Threes(dice)
	scores["Fours"] = Fours(dice)
	scores["Fives"] = Fives(dice)
	scores["Sixes"] = Sixes(dice)
	scores["OnePair"] = Pair(dice)
	scores["TwoPairs"] = TwoPairs(dice)
	scores["ThreeOfAKind"] = Toak(dice)
	scores["FourOfAKind"] = Foak(dice)
	scores["SmallStraight"] = SmallStraight(dice)
	scores["LargeStraight"] = LargeStraight(dice)
	scores["FullHouse"] = FullHouse(dice)
	scores["Chance"] = Chance(dice)
	scores["Yatzy"] = Yatzy(dice)
	scores["Bonus"] = Bonus(currentPlayerScore)
	scores["Total"] = Total(currentPlayerScore)
	return scores
}
