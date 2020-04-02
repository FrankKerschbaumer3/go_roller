package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	math_rand "math/rand"

	"github.com/FrankKerschbaumer3/go_roller/roll"
)

// Creates seed to have random dice rolls
func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

//Function that will accept ints and booleans to output a dice total.
func main() {

	//would call function to create gui
	//gui would take in user inputs and return inputs/flags here

	diceToRoll := flag.Int("dicetoroll", 1, "What dice will you roll?")
	amountOfDice := flag.Int("amountofdice", 1, "How many dice will you roll?")
	atAdvantage := flag.Bool("advantage", false, "Roll with Advantage?")
	atDisadvantage := flag.Bool("disadvantage", false, "Roll with Disadvantage?")
	modifier := flag.Int("modifier", 0, "What is your modifier?")

	flag.Parse()

	sum := roll.Roll(*diceToRoll, *amountOfDice, *modifier, *atAdvantage, *atDisadvantage)

	fmt.Println(sum)

	//sum would then get passed to the gui as an output for what was rolled
	//would then loop back and wait for more inputs in the gui
}
