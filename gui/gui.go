package gui

import (
	"strconv"

	"github.com/FrankKerschbaumer3/go_roller/roll"
	"github.com/rivo/tview"
)

// Gui creates the cli gui for the application
func Gui() {
	app := tview.NewApplication().EnableMouse(true)
	diceToRoll := tview.NewDropDown().
		SetLabel("Dice to Roll").
		SetOptions([]string{"20", "12", "10", "8", "6", "4", "100"}, nil)
	amountOfDice := tview.NewInputField().
		SetLabel("How many to roll?").
		SetFieldWidth(2).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(nil)
	modifier := tview.NewInputField().
		SetLabel("Modifier?").
		SetFieldWidth(2).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(nil)
	rolledAmount := tview.NewInputField().
		SetText("").
		SetLabel("").
		SetFieldWidth(-1).
		SetAcceptanceFunc(nil).
		SetDoneFunc(nil)
	form := tview.NewForm().
		AddFormItem(diceToRoll).
		AddFormItem(amountOfDice).
		AddFormItem(modifier).
		AddButton("Roll", func() {
			_, rolling := diceToRoll.GetCurrentOption()
			howMany := amountOfDice.GetText()
			modf := modifier.GetText()
			dice, amount, mod := strIn(rolling, howMany, modf)
			roll := roll.Roll(dice, amount, mod, false, false)
			sum := intOut(roll)
			rolledAmount.
				SetText(sum).
				SetLabel("Roll:").
				SetFieldWidth(3)
		}).
		AddButton("Advantage", func() {
			_, rolling := diceToRoll.GetCurrentOption()
			howMany := "2"
			modf := modifier.GetText()
			dice, amount, mod := strIn(rolling, howMany, modf)
			roll := roll.Roll(dice, amount, mod, true, false)
			sum := intOut(roll)
			rolledAmount.
				SetText(sum).
				SetLabel("Roll:").
				SetFieldWidth(3)
		}).
		AddButton("Disadvantage", func() {
			_, rolling := diceToRoll.GetCurrentOption()
			howMany := "2"
			modf := modifier.GetText()
			dice, amount, mod := strIn(rolling, howMany, modf)
			roll := roll.Roll(dice, amount, mod, false, true)
			sum := intOut(roll)
			rolledAmount.
				SetText(sum).
				SetLabel("Roll:").
				SetFieldWidth(3)
		}).
		AddButton("Reset", func() {
			diceToRoll.SetCurrentOption(-1)
			amountOfDice.SetText("")
			modifier.SetText("")
			rolledAmount.SetLabel("").SetText("").SetFieldWidth(-1)
		}).
		AddButton("Quit", func() {
			app.Stop()
		}).
		AddFormItem(rolledAmount)
	form.SetBorder(true).SetTitle("Go_Roller").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}

func strIn(a, b, c string) (dice, amount, mod int) {
	dice, _ = strconv.Atoi(a)
	amount, _ = strconv.Atoi(b)
	mod, _ = strconv.Atoi(c)
	return
}

func intOut(num int) string {
	i := strconv.Itoa(num)
	return i
}
