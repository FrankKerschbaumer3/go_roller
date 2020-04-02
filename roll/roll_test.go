package roll

import (
	"math/rand"
	"testing"
)

var roll1 = []struct {
	dice         int
	amount       int
	modifier     int
	advantage    bool
	disadvantage bool
	expected     int
}{
	{20, 1, 5, false, false, 9},
	{20, 2, 10, true, false, 19},
	{20, 2, -3, false, true, 1},
	{8, 8, 6, false, false, 41},
	{12, 3, -5, false, false, 17},
	{20, 1, 0, false, false, 4},
	{6, 0, 5, false, false, 5},
}

var roll2 = []struct {
	dice         int
	amount       int
	modifier     int
	advantage    bool
	disadvantage bool
	expected     int
}{
	{20, 1, 9, false, false, 22},
	{20, 2, -4, true, false, 9},
	{20, 2, 8, false, true, 12},
	{10, 10, 3, false, false, 67},
	{12, 4, -2, false, false, 17},
	{20, 1, 0, false, false, 13},
	{6, 0, 10, false, false, 10},
}

func TestRoll1(t *testing.T) {
	for _, roll := range roll1 {
		rand.Seed(100)
		t.Run("Roll", func(t *testing.T) {
			result := Roll(roll.dice, roll.amount, roll.modifier, roll.advantage, roll.disadvantage)
			if result != roll.expected {
				t.Error("Something went wrong :(")
			}
		})
	}
}

func TestRoll2(t *testing.T) {
	for _, roll := range roll2 {
		rand.Seed(666)
		t.Run("Roll", func(t *testing.T) {
			result := Roll(roll.dice, roll.amount, roll.modifier, roll.advantage, roll.disadvantage)
			if result != roll.expected {
				t.Error("Something went wrong :(")
			}
		})
	}
}
