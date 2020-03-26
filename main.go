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

	slice := make([]int, 0, *amountOfDice)
	rand.Seed(time.Now().UnixNano())

	if *atAdvantage == true {

		for i := 1; i <= *amountOfDice; i++ {
			num := 1 + rand.Intn(*diceToRoll)
			slice = append(slice, num)
		}
		max := slice[0]

		for i := 1; i < *amountOfDice; i++ {

			if max < slice[i] {

				max = slice[i]
			}
		}
		fmt.Println(*modifier + max)
	}

	if *atDisadvantage == true {
		for i := 1; i <= *amountOfDice; i++ {
			num := 1 + rand.Intn(*diceToRoll)
			slice = append(slice, num)
		}
		min := slice[0]

		for i := 1; i < *amountOfDice; i++ {

			if min > slice[i] {

				min = slice[i]
			}
		}
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
