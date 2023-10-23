// main_test.go
package main

import (
	"testing"
	//"github.com/fatih/color"
)

func TestColorizeText(t *testing.T) {
	//c := color.New(color.FgCyan)
	coloredText := colorizeText("This text should be cyan!")

	expected := "\x1b[36mThis text should be cyan!\x1b[0m"

	if coloredText != expected {
		t.Errorf("Expected %s but got %s", expected, coloredText)
	}
}
func TestMetricsMain(t *testing.T) {
}
