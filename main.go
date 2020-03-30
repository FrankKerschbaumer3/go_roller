package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

//Function to roll dice
func main() {

	//Flags to parse dice conditions
	diceToRoll := flag.Int("dicetoroll", 1, "What dice will you roll?")
	amountOfDice := flag.Int("amountofdice", 1, "How many dice will you roll?")
	atAdvantage := flag.Bool("advantage", false, "Roll with Advantage?")
	atDisadvantage := flag.Bool("disadvantage", false, "Roll with Disadvantage?")
	modifier := flag.Int("modifier", 0, "What is your modifier?")

	flag.Parse()

	slice := make([]int, 0)

	rand.Seed(time.Now().UnixNano())

	if *atAdvantage == true {
		for i := 1; i <= *amountOfDice; i++ {
			num := 1 + rand.Intn(*diceToRoll)
			slice = append(slice, num)
		}
		// Pass ints to max function
		max := Max(slice)

		sum := *modifier + max

		fmt.Println(sum)
	}

	if *atDisadvantage == true {
		for i := 1; i <= *amountOfDice; i++ {
			num := 1 + rand.Intn(*diceToRoll)
			slice = append(slice, num)
		}
		// Pass ints to min function
		min := Min(slice)

		sum := *modifier + min

		fmt.Println(sum)
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
