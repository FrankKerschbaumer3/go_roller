package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// MinMax creates a min and max value.
func MinMax(a []int) (int, int) {
	min := a[0]
	max := a[0]
	for _, value := range a {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

//Function to roll dice
func main() {

	//Flags to parse dice conditions
	diceToRoll := flag.Int("dicetoroll", 1, "What dice will you roll?")
	amountOfDice := flag.Int("amountofdice", 1, "How many dice will you roll?")
	atAdvantage := flag.Bool("advantage", false, "Roll with Advantage?")
	atDisadvantage := flag.Bool("disadvantage", false, "Roll with Disadvantage?")
	modifier := flag.Int("modifier", 0, "What is your modifier?")
	flag.Parse()

	slice := make([]int, 0, *amountOfDice)
	rand.Seed(time.Now().UnixNano())

	if *atAdvantage == true {
		for i := 1; i <= *amountOfDice; i++ {
			num := 1 + rand.Intn(*diceToRoll)
			slice = append(slice, num)
		}
		_, max := MinMax(slice)

		fmt.Println(*modifier + max)
	}

	if *atDisadvantage == true {
		for i := 1; i <= *amountOfDice; i++ {
			num := 1 + rand.Intn(*diceToRoll)
			slice = append(slice, num)
		}

		min, _ := MinMax(slice)
		fmt.Println(*modifier + min)
	}

	if *atAdvantage == false && *atDisadvantage == false {
		sum := 0
		for i := 1; i <= *amountOfDice; i++ {
			num := 1 + rand.Intn(*diceToRoll)
			sum += num
		}
		fmt.Println(*modifier + sum)
	}
}
