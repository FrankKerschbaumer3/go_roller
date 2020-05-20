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
	Sides    int   `json:"sides" binding:"required"`
	Rolled   []int `json:"rolled" binding:"required"`
	Modifier int   `json:"modifier"`
	Total    int   `json:"total"`
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
			Sides:    out.Sides,
			Rolled:   out.Rolled,
			Modifier: out.Modifier,
			Total:    out.Total,
		}
		rollOutputs = append(rollOutputs, output)
	}

	c.JSON(http.StatusOK, rollOutputs)
}

// Modifier structure for advnatage and disadvantage rolls
type Modifier struct {
	Modifier int `json:"modifier"`
}

// AdvantageRoll passes the json to the roll function
func AdvantageRoll(c *gin.Context) {

	var input Modifier
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rollInput := roll.Input{
		Sides:    20,
		Amount:   1,
		Modifier: input.Modifier,
	}

	advantageOutput := roll.Roll(rollInput, true, false)

	output := Output{
		Sides:    advantageOutput.Sides,
		Rolled:   advantageOutput.Rolled,
		Modifier: advantageOutput.Modifier,
		Total:    advantageOutput.Total,
	}

	c.JSON(http.StatusOK, output)
}

// DisadvantageRoll passes the json to the roll function
func DisadvantageRoll(c *gin.Context) {

	var input Modifier
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rollInput := roll.Input{
		Sides:    20,
		Amount:   1,
		Modifier: input.Modifier,
	}

	disadvantageOutput := roll.Roll(rollInput, true, false)

	output := Output{
		Sides:    disadvantageOutput.Sides,
		Rolled:   disadvantageOutput.Rolled,
		Modifier: disadvantageOutput.Modifier,
		Total:    disadvantageOutput.Total,
	}

	c.JSON(http.StatusOK, output)
}

// Pong responds from ping
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
