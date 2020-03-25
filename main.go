package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	diceToRoll := flag.Int("dicetoroll", 1, "What dice will you roll?")
	amountOfDice := flag.Int("amountofdice", 1, "How many dice will you roll?")
	atAdvantage := flag.Bool("advantage", false, "Roll with Advantage?")
	atDisadvantage := flag.Bool("disadvantage", false, "Roll with Disdvantage?")
	modifier := flag.Int("modifier", 0, "What is your modifier?")
	flag.Parse()

	sum := 0

	rand.Seed(time.Now().UnixNano())

	if *atAdvantage == true {
		fmt.Println("Yo")
	}

	if *atDisadvantage == true {
		fmt.Println("Yo")
	}

	if *atAdvantage == false && *atDisadvantage == false {
		for i := 1; i <= *amountOfDice; i++ {
			num := 1 + rand.Intn(*diceToRoll)
			sum += num
		}
		fmt.Println(*modifier + sum)
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
