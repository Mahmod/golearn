// main.go
package main

import (
	"fmt"
	"github.com/fatih/color"
)

func colorizeText(text string) string {
	c := color.New(color.FgCyan)
	return c.Sprint(text)
}

func main() {
	fmt.Println("Hello, Go with dependencies!")
}
