package roll

import (
	"fmt"
	"math/rand"

	"github.com/FrankKerschbaumer3/go_roller/minmax"
)

// Input struct for what to roll
type Input struct {
	Sides    int
	Amount   int // Amount to roll
	Modifier int
}

// Result struct for the return values
type Result struct {
	Sides  int
	Rolled []int
	Total  int
}

// Roll will take in arguments and output what was rolled as the first returned output
// and the total on the second returned output
func Roll(input Input, atAdvantage, atDisadvantage bool) *Result {

	slice := make([]int, 0)

	if atAdvantage == true {
		input.Amount = 2
		for i := 1; i <= input.Amount; i++ {
			num := 1 + rand.Intn(input.Sides)
			slice = append(slice, num)
		}
		// Pass ints to Max minmax.go function

		max := minmax.Max(slice)
		return &Result{
			Sides:  input.Amount,
			Rolled: slice,
			Total:  input.Modifier + max,
		}
	} else if atDisadvantage == true {
		input.Amount = 2
		for i := 1; i <= input.Amount; i++ {
			num := 1 + rand.Intn(input.Sides)
			slice = append(slice, num)
		}

		// Pass ints to Min minmax.go function
		min := minmax.Min(slice)
		return &Result{
			Sides:  input.Amount,
			Rolled: slice,
			Total:  input.Modifier + min,
		}
	} else {
		sum := 0
		num := 0
		for i := 1; i <= input.Amount; i++ {
			num = 1 + rand.Intn(input.Sides)
			slice = append(slice, num)
			sum += num
		}
		return &Result{
			Sides:  input.Amount,
			Rolled: slice,
			Total:  input.Modifier + sum,
		}
	}
}

// All lkjshdalfj
func All(inputs []Input) []*Result {

	slice := make([]*Result, 0)

	for _, input := range inputs {
		rolls := Roll(input, false, false)
		slice = append(slice, rolls)
		fmt.Println(rolls)
	}
	return slice
}

// call roll multiple times with the various dice to roll
