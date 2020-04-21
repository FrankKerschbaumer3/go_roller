package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"

	"github.com/FrankKerschbaumer3/go_roller/controllers"
	"github.com/gin-gonic/gin"
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

func main() {
	r := gin.Default()
	r.GET("/ping", controllers.Pong)
	r.POST("/roll", controllers.NormalRoll)
	r.POST("/roll/advantage", controllers.AdvantageRoll)
	r.POST("/roll/disadvantage", controllers.DisadvantageRoll)
	r.Run(":8080")
}
