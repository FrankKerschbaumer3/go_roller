package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/FrankKerschbaumer3/go_roller/roll"
)

//Function that will accept ints and booleans to output a dice total.
func main() {

	//Flags to parse dice conditions
	diceToRoll := flag.Int("dicetoroll", 1, "What dice will you roll?")
	amountOfDice := flag.Int("amountofdice", 1, "How many dice will you roll?")
	atAdvantage := flag.Bool("advantage", false, "Roll with Advantage?")
	atDisadvantage := flag.Bool("disadvantage", false, "Roll with Disadvantage?")
	modifier := flag.Int("modifier", 0, "What is your modifier?")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	sum := roll.Roll(*diceToRoll, *amountOfDice, *modifier, *atAdvantage, *atDisadvantage)

	fmt.Println(sum)
}
