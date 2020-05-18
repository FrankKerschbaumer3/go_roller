package controllers

import (
	"net/http"

	"github.com/FrankKerschbaumer3/go_roller/roll"
	"github.com/gin-gonic/gin"
)

// Request json structure
type Request struct {
	Roll []Input `json:"roll" binding:"required"`
}

// Input json structure
type Input struct {
	Sides    int `json:"sides" binding:"required"`
	Amount   int `json:"amount" binding:"required"`
	Modifier int `json:"modifier"`
}

// Output Json Structure
type Output struct {
	Sides  int   `json:"sides" binding:"required"`
	Rolled []int `json:"rolled" binding:"required"`
	Total  int   `json:"total"`
}

//NormalRoll passes the json to the roll function
func NormalRoll(c *gin.Context) {

	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rollInputs := make([]roll.Input, 0)

	for _, in := range request.Roll {
		input := roll.Input(in)
		rollInputs = append(rollInputs, input)
	}

	rolling := roll.All(rollInputs)

	rollOutputs := make([]Output, 0)

	for _, out := range rolling {
		output := Output{
			Sides:  out.Sides,
			Rolled: out.Rolled,
			Total:  out.Total,
		}
		rollOutputs = append(rollOutputs, output)
	}

	c.JSON(http.StatusOK, rollOutputs)
}

// AdvantageRoll passes the json to the roll function
// func AdvantageRoll(c *gin.Context) {

// 	var input Input
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	input = Input{
// 		Sides:    input.Sides,
// 		Amount:   input.Amount,
// 		Modifier: input.Modifier,
// 	}

// 	advantageOutput := roll.Roll(roll.Input{}, true, false)

// 	output := Output{
// 		Sides:  advantageOutput.Sides,
// 		Rolled: advantageOutput.Rolled,
// 		Total:  advantageOutput.Total,
// 	}

// 	c.JSON(http.StatusOK, output)
// }

// // DisadvantageRoll passes the json to the roll function
// func DisadvantageRoll(c *gin.Context) {

// 	var input Input
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	input = Input{
// 		Sides:    input.Sides,
// 		Amount:   input.Amount,
// 		Modifier: input.Modifier,
// 	}

// 	disadvantageOutput := roll.Roll(roll.Input{}, false, true)

// 	rollOutput := Output{
// 		Sides:  disadvantageOutput.Sides,
// 		Rolled: disadvantageOutput.Rolled,
// 		Total:  disadvantageOutput.Total,
// 	}

// 	c.JSON(http.StatusOK, rollOutput)
// }

// Pong responds from ping
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
