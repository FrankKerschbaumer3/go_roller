package minmax

import (
	"testing"
)

var minrolls = []struct {
	in       []int
	expected int
}{
	{[]int{1, 13}, 1},
	{[]int{7, 3}, 3},
	{[]int{4, 26}, 4},
	{[]int{18, 12, 22, 7}, 7},
	{[]int{3}, 3},
}

var maxrolls = []struct {
	in       []int
	expected int
}{
	{[]int{1, 13}, 13},
	{[]int{7, 3}, 7},
	{[]int{4, 26}, 26},
	{[]int{18, 12, 22, 7}, 22},
	{[]int{3}, 3},
}

func TestMin(t *testing.T) {
	for _, rolls := range minrolls {
		t.Run("Min", func(t *testing.T) {
			result := Min(rolls.in)
			if result != rolls.expected {
				t.Error("Min fuction did not return expected value")
			}
		})
	}
}

func TestMax(t *testing.T) {
	for _, rolls := range maxrolls {
		t.Run("Max", func(t *testing.T) {
			result := Max(rolls.in)
			if result != rolls.expected {
				t.Error("Max fuction did not return expected value")
			}
		})
	}
}
