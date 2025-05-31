package tui

import (
	"log"

	"github.com/Leathal1/greycli/internal/views"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Run starts the full-screen terminal UI
func Run() error {
	app := tview.NewApplication()

	// Help bar
	helpTextView := tview.NewTextView().
		SetText("[J] Jobs  [M] Models  [C] Credits  [S] Submit  [Q] Quit").
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorWhite).
		SetDynamicColors(true).
		SetChangedFunc(func() { app.Draw() }).
		SetBorder(true).
		SetTitle("Help").
		SetBackgroundColor(tcell.ColorBlack)

	helpVisible := true

	// Views
	jobsView := views.NewJobsView()
	jobsView.SetBackgroundColor(tcell.ColorGray)

	var modelsView *tview.List = views.NewModelsView()
	modelsView.SetBackgroundColor(tcell.ColorGray)

	creditsView := views.NewCreditsView()
	creditsView.SetBackgroundColor(tcell.ColorGray)

	submitView := views.NewSubmitView()
	submitView.SetBackgroundColor(tcell.ColorGray)
	submitView.SetTitle("Submit – Selected: none")

	pages := tview.NewPages().
		AddPage("jobs", jobsView, true, true).
		AddPage("models", modelsView, true, false).
		AddPage("credits", creditsView, true, false).
		AddPage("submit", submitView, true, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(helpTextView, 1, 0, false)

	// Help toggle function
	toggleHelp := func() {
		if helpVisible {
			layout.RemoveItem(helpTextView)
		} else {
			layout.AddItem(helpTextView, 1, 0, false)
		}
		helpVisible = !helpVisible
		app.Draw()
	}

	// Global key bindings
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q', 'Q':
			app.Stop()
		case 'j', 'J':
			pages.SwitchToPage("jobs")
		case 'm', 'M':
			pages.SwitchToPage("models")
		case 'c', 'C':
			pages.SwitchToPage("credits")
		case 's', 'S':
			pages.SwitchToPage("submit")
		case 'h', 'H':
			toggleHelp()
		}
		return event
	})

	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		log.Printf("TUI error: %v", err)
		return err
	}

	return nil
}
