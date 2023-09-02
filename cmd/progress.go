package cmd

import (
	"Progress/utils"
	"fmt"
	"strings"
)

const (
	emptyChar  = "▒"
	filledChar = "█"
)

type ProgressBar struct {
	MaxValue           float64
	Width              float64
	currentValue       float64
	currentProgressBar string
}

func NewProgressBar(maxValue, width float64) ProgressBar {
	utils.GetTerminalSize()

	if float64(utils.Window.Width) < width {
		width = float64(utils.Window.Width) * 0.90
	}
	return ProgressBar{
		MaxValue: maxValue,
		Width:    width,
	}
}

func (bar *ProgressBar) Update(value float64) error {
	if value < bar.currentValue {
		return fmt.Errorf("update value is less than current value")
	}
	bar.currentValue = value
	bar.renderBar()
	return nil
}

func (bar *ProgressBar) renderBar() {

	bar.currentProgressBar = ""
	bar.CalculateBar()
	filledCount := (bar.currentValue / bar.MaxValue) * bar.Width
	percent := int(filledCount / bar.Width * 100)

	fmt.Print("\x1B[?47h") // Save screen
	utils.ClearTerminal()
	fmt.Print("\x1B[?47l")                            // Restore Screen
	fmt.Printf("\x1B[%d;%dH", utils.Window.Height, 0) // moving cursor down
	fmt.Printf("%v : %v%v", bar.currentProgressBar, percent, "%")
}

func (bar *ProgressBar) CalculateBar() {
	ratio := bar.currentValue / bar.MaxValue
	filledCount := ratio * bar.Width
	emptyCount := bar.Width - filledCount
	bar.currentProgressBar = fmt.Sprintf("%s%s", strings.Repeat(filledChar, int(filledCount)), strings.Repeat(emptyChar, int(emptyCount)))
}

func (bar *ProgressBar) CleanUp() {
	fmt.Printf("\x1B[%d;%dH", utils.Window.Height, 0) // moving cursor down
	fmt.Print("\x1B[0K")
	fmt.Print("\x1B8")
}
