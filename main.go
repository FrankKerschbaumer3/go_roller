package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"

	"github.com/FrankKerschbaumer3/go_roller/gui"
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

	gui.Gui()

}
