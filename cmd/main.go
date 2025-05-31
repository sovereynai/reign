package main

import (
	"log"

	"github.com/Leathal1/greycli/internal/api"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Leathal1/greycli/internal/views"
)

func main() {
	app := tview.NewApplication()

	helpTextView := tview.NewTextView()
	helpTextView.SetText("[J] Jobs  |  [M] Models  |  [C] Credits  |  [S] Submit  |  [Q] Quit  |  [H] Toggle Help")
	helpTextView.SetTextAlign(tview.AlignCenter)
	helpTextView.SetTextColor(tcell.ColorWhite)
	helpTextView.SetBackgroundColor(tcell.ColorBlack)
	helpTextView.SetBorder(false)

	helpVisible := true
	var toggleHelp func()

	pages := tview.NewPages()

	// Initialize views
	var selectedModel string
	jobsView := views.NewJobsView()
	var modelsView *tview.List = views.NewModelsView()
	creditsView := views.NewCreditsView()
	submitView := views.NewSubmitView()
	submitView.SetTitle("Submit – Selected: none")

	// Set initial focus to the model selector
	reloadModels := func() {
		modelsView.Clear()
		models, err := api.FetchModels()
		if err != nil {
			modelsView.AddItem("Error loading models", "", 0, nil)
			return
		}
		for _, m := range models {
			mCopy := m
			modelsView.AddItem(mCopy, "", 0, func() {
				selectedModel = mCopy
				app.QueueUpdateDraw(func() {
					submitView.SetTitle("Submit – Selected: " + selectedModel)
				})
			})
		}
	}
	app.SetFocus(modelsView)

	// Add each view as a page
	pages.AddPage("jobs", jobsView, true, true)
	pages.AddPage("models", modelsView, true, false)
	pages.AddPage("credits", creditsView, true, false)
	pages.AddPage("submit", submitView, true, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(helpTextView, 1, 0, false)

	toggleHelp = func() {
		if helpVisible {
			layout.RemoveItem(helpTextView)
			helpVisible = false
		} else {
			layout.AddItem(helpTextView, 1, 0, false)
			helpVisible = true
		}
	}

	// Set global keybindings
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
		case 'r', 'R':
			reloadModels()
		}
		return event
	})

	// Set theme colors
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack
	tview.Styles.ContrastBackgroundColor = tcell.ColorGray
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorGray
	tview.Styles.BorderColor = tcell.ColorWhite
	tview.Styles.TitleColor = tcell.ColorWhite
	tview.Styles.GraphicsColor = tcell.ColorGreen
	tview.Styles.PrimaryTextColor = tcell.ColorWhite
	tview.Styles.TertiaryTextColor = tcell.ColorGray

	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		log.Fatalf("Error running TUI: %v", err)
	}
}
