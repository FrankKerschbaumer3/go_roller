package roll

import (
	"math/rand"
	"testing"
)

var tests = []struct {
	input        Input
	advantage    bool
	disadvantage bool
	sum          int
}{
	{Input{Sides: 20, Amount: 1, Modifier: 5}, false, false, 9},
	{Input{Sides: 8, Amount: 8, Modifier: 6}, false, false, 41},
	{Input{Sides: 20, Amount: 2, Modifier: 10}, true, false, 19},
	{Input{Sides: 20, Amount: 2, Modifier: -3}, false, true, 1},
	{Input{Sides: 12, Amount: 3, Modifier: -5}, false, false, 17},
	{Input{Sides: 20, Amount: 1, Modifier: 0}, false, false, 4},
	{Input{Sides: 6, Amount: 0, Modifier: 5}, false, false, 5},
}

func TestRoll(t *testing.T) {
	for _, expected := range tests {
		rand.Seed(100)
		t.Run("Roll", func(t *testing.T) {
			result := Roll(expected.input, expected.advantage, expected.disadvantage)
			if result.Total != expected.sum {
				t.Error("expected", expected.sum, "but got", result.Total, ".", "Rolled", expected.input.Amount, "d", expected.input.Sides, result.Rolled, "with a modifier", expected.input.Modifier)
			}
		})
	}
}

var test2 = []struct {
	input        Input
	advantage    bool
	disadvantage bool
	sum          int
}{
	{Input{Sides: 20, Amount: 1, Modifier: 5}, false, false, 9},
	{Input{Sides: 8, Amount: 8, Modifier: 6}, false, false, 34},
}

func TestAll(t *testing.T) {
	rand.Seed(100)
	t.Run("Roll", func(t *testing.T) {
		inputs := []Input{test2[0].input, test2[1].input}
		result := All(inputs)
		if result[0].Total != test2[0].sum {
			expected := test2[0]
			result := result[0]
			t.Error("Roll1 expected", expected.sum, "but got", result.Total, ".", "Rolled", expected.input.Amount, "d", expected.input.Sides, result.Rolled, "with a modifier", expected.input.Modifier)
		}
		if result[1].Total != test2[1].sum {
			expected := test2[1]
			result := result[1]
			t.Error("Roll2 expected", expected.sum, "but got", result.Total, ".", "Rolled", expected.input.Amount, "d", expected.input.Sides, result.Rolled, "with a modifier", expected.input.Modifier)
		}
	})
}
