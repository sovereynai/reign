// cmd/main.go
package main

import (
	"log"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	text := tview.NewTextView().
		SetText("Welcome to GreyCLI — Fist AI Network").
		SetTextAlign(tview.AlignCenter)

	if err := app.SetRoot(text, true).EnableMouse(true).Run(); err != nil {
		log.Fatalf("Error running TUI: %v", err)
	}
}
