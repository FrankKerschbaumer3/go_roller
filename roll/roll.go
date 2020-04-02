package roll

import (
	"math/rand"

	"github.com/FrankKerschbaumer3/go_roller/minmax"
)

//Roll will take in arguments and output dice totals based on arguments
func Roll(diceToRoll, amountOfDice, modifier int, atAdvantage, atDisadvantage bool) int {

	slice := make([]int, 0)

	if atAdvantage == true {
		for i := 1; i <= amountOfDice; i++ {
			num := 1 + rand.Intn(diceToRoll)
			slice = append(slice, num)
		}
		// Pass ints to Max minmax.go function
		max := minmax.Max(slice)
		return modifier + max
	} else if atDisadvantage == true {
		for i := 1; i <= amountOfDice; i++ {
			num := 1 + rand.Intn(diceToRoll)
			slice = append(slice, num)
		}
		// Pass ints to Min minmax.go function
		min := minmax.Min(slice)
		return modifier + min
	} else {
		sum := 0
		num := 0
		for i := 1; i <= amountOfDice; i++ {
			num = 1 + rand.Intn(diceToRoll)
			sum += num
		}
		return modifier + sum
	}
}
