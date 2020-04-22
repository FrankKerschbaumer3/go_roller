package controllers

import (
	"net/http"

	"github.com/FrankKerschbaumer3/go_roller/roll"
	"github.com/gin-gonic/gin"
)

// Input json structure
type Input struct {
	Name string `json:"name" binding:"required"`
	Dice int    `json:"dice" binding:"required"`
	Num  int    `json:"num"`
	Mod  int    `json:"mod"`
}

// Output Json Structure
type Output struct {
	Name   string `json:"name" binding:"required"`
	Rolled []int  `json:"dice" binding:"required"`
	Mod    int    `json:"mod" binding:"required"`
	Total  int    `json:"total" binding:"required"`
}

//NormalRoll passes the json to the roll function
func NormalRoll(c *gin.Context) {

	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input = Input{
		Name: input.Name,
		Dice: input.Dice,
		Num:  input.Num,
		Mod:  input.Mod}

	rolled := roll.Roll(input.Dice, input.Num, input.Mod, false, false)

	output := Output{
		Name:   input.Name,
		Rolled: rolled.Rolled,
		Mod:    input.Mod,
		Total:  rolled.Sum,
	}

	c.JSON(http.StatusOK, output)
}

//AdvantageRoll passes the json to the roll function
func AdvantageRoll(c *gin.Context) {

	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input = Input{
		Name: input.Name,
		Dice: input.Dice,
		Num:  input.Num,
		Mod:  input.Mod}

	rolled := roll.Roll(input.Dice, input.Num, input.Mod, true, false)

	output := Output{
		Name:   input.Name,
		Rolled: rolled.Rolled,
		Mod:    input.Mod,
		Total:  rolled.Sum,
	}

	c.JSON(http.StatusOK, output)
}

//DisadvantageRoll passes the json to the roll function
func DisadvantageRoll(c *gin.Context) {

	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input = Input{
		Name: input.Name,
		Dice: input.Dice,
		Num:  input.Num,
		Mod:  input.Mod}

	rolled := roll.Roll(input.Dice, input.Num, input.Mod, false, true)

	output := Output{
		Name:   input.Name,
		Rolled: rolled.Rolled,
		Mod:    input.Mod,
		Total:  rolled.Sum,
	}

	c.JSON(http.StatusOK, output)
}

// Pong responds from ping
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
