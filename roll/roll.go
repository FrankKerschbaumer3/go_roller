package roll

import (
	"math/rand"

	"github.com/FrankKerschbaumer3/go_roller/minmax"
)

// Result struct for the return values
type Result struct {
	Rolled []int
	Sum    int
}

// Roll will take in arguments and output what was rolled as the first returned output
// and the total on the second returned output
func Roll(diceToRoll, amountOfDice, modifier int, atAdvantage, atDisadvantage bool) *Result {

	slice := make([]int, 0)

	if atAdvantage == true {
		amountOfDice = 2
		for i := 1; i <= amountOfDice; i++ {
			num := 1 + rand.Intn(diceToRoll)
			slice = append(slice, num)
		}
		// Pass ints to Max minmax.go function

		max := minmax.Max(slice)
		return &Result{
			Rolled: slice,
			Sum:    modifier + max,
		}
	} else if atDisadvantage == true {
		amountOfDice = 2
		for i := 1; i <= amountOfDice; i++ {
			num := 1 + rand.Intn(diceToRoll)
			slice = append(slice, num)
		}

		// Pass ints to Min minmax.go function
		min := minmax.Min(slice)
		return &Result{
			Rolled: slice,
			Sum:    modifier + min,
		}
	} else {
		sum := 0
		num := 0
		for i := 1; i <= amountOfDice; i++ {
			num = 1 + rand.Intn(diceToRoll)
			sum += num
		}
		return &Result{
			Rolled: []int{sum},
			Sum:    modifier + sum,
		}
	}
}
