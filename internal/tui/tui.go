package tui

import (
   "fmt"
   "time"

   "github.com/Leathal1/greycli/internal/api"
   "github.com/gdamore/tcell/v2"
   "github.com/rivo/tview"
)

func NewModelsView() *tview.List {
	modelList := tview.NewList()
	modelList.ShowSecondaryText(false)
   // Styling
   modelList.SetBorder(true).
       SetBorderColor(tcell.ColorWhite).
       SetTitle("Models").
       SetTitleColor(tcell.ColorWhite)
   modelList.SetMainTextColor(tcell.ColorWhite)
   modelList.SetSelectedTextColor(tcell.ColorBlack)
   modelList.SetSelectedBackgroundColor(tcell.ColorLime)
   modelList.SetHighlightFullLine(true)

	models, err := api.FetchModels()
	if err != nil {
		modelList.AddItem("Failed to load models", "", 0, nil)
	} else {
		for _, model := range models {
			m := model // capture loop var
			modelList.AddItem(m, "", 0, func() {
				fmt.Printf("Selected model: %s\n", m)
			})
		}
	}

	return modelList
}

// NewCreditsView creates a TextView with a loading spinner while fetching credits.
func NewCreditsView(app *tview.Application) *tview.TextView {
   creditsView := tview.NewTextView()
   creditsView.SetBorder(true).
       SetBorderColor(tcell.ColorWhite).
       SetTitle("Credits").
       SetTitleColor(tcell.ColorWhite)
   creditsView.SetTextColor(tcell.ColorWhite)

   // Spinner setup
   spinner := []rune{'|', '/', '-', '\\'}
   idx := 0
   ticker := time.NewTicker(200 * time.Millisecond)
   done := make(chan struct{})
   // Initial text
   creditsView.SetText(fmt.Sprintf("Credits: loading... %c", spinner[0]))

   // Animate spinner
   go func() {
       for {
           select {
           case <-done:
               return
           case <-ticker.C:
               idx++
               ch := spinner[idx%len(spinner)]
               app.QueueUpdateDraw(func() {
                   creditsView.SetText(fmt.Sprintf("Credits: loading... %c", ch))
               })
           }
       }
   }()

   // Fetch credits and stop spinner
   go func() {
       credits, err := api.FetchCredits()
       ticker.Stop()
       close(done)
       app.QueueUpdateDraw(func() {
           if err != nil {
               creditsView.SetText("Credits: error")
           } else {
               creditsView.SetText(fmt.Sprintf("Credits: %d", credits))
           }
       })
   }()

   return creditsView
}

func NewJobsView() *tview.Table {
   table := tview.NewTable()
   table.SetBorders(false)
   table.SetBorder(true).
       SetBorderColor(tcell.ColorWhite).
       SetTitle("Jobs").
       SetTitleColor(tcell.ColorWhite)
   // Make selectable and header fixed
   table.SetSelectable(true, false)
   table.SetFixed(1, 0)
   table.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorLime).Foreground(tcell.ColorBlack))

   // Header row
   headerColor := tcell.ColorLime
   table.SetCell(0, 0, tview.NewTableCell("ID").SetTextColor(headerColor).SetSelectable(false))
   table.SetCell(0, 1, tview.NewTableCell("Model").SetTextColor(headerColor).SetSelectable(false))
   table.SetCell(0, 2, tview.NewTableCell("Status").SetTextColor(headerColor).SetSelectable(false))

   // Fetch jobs (errors are silently ignored, showing header only)
   jobs, err := api.FetchJobs()
   if err != nil {
       return table
   }

   // Populate rows
   for i, job := range jobs {
       row := i + 1
       table.SetCell(row, 0, tview.NewTableCell(job.ID).SetTextColor(tcell.ColorWhite))
       table.SetCell(row, 1, tview.NewTableCell(job.Model).SetTextColor(tcell.ColorWhite))
       table.SetCell(row, 2, tview.NewTableCell(job.Status).SetTextColor(tcell.ColorWhite))
   }
   return table
}

// Run launches the interactive terminal UI application.
// Run launches the interactive terminal UI application, showing one view at a time
// with a help bar for switching between Credits, Models, and Jobs.
func Run() error {
   // Set a uniform grey background for all primitives
   tview.Styles.PrimitiveBackgroundColor = tcell.ColorGrey
   app := tview.NewApplication()

   // Create individual views, passing the app for animations
   creditsView := NewCreditsView(app)
   modelsView := NewModelsView()
   jobsView := NewJobsView()

   // Page container: show one view at a time
   pages := tview.NewPages().
       AddPage("credits", creditsView, true, true).
       AddPage("models", modelsView, true, false).
       AddPage("jobs", jobsView, true, false)

   // Help bar with keybindings
   // Help bar with keybindings (static text, no dynamic colors)
   helpBar := tview.NewTextView()
   helpBar.SetText(" [C]redits  [M]odels  [J]obs  [Q]uit ")
   helpBar.SetTextAlign(tview.AlignCenter)
   helpBar.SetTextColor(tcell.ColorWhite)
   helpBar.SetBackgroundColor(tcell.ColorDarkGrey)

   // Layout: pages on top, help bar at bottom
   layout := tview.NewFlex().SetDirection(tview.FlexRow).
       AddItem(pages, 0, 1, true).
       AddItem(helpBar, 1, 0, false)

   // Global key handling to switch pages
   app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
       switch event.Rune() {
       case 'c', 'C':
           pages.SwitchToPage("credits")
       case 'm', 'M':
           pages.SwitchToPage("models")
       case 'j', 'J':
           pages.SwitchToPage("jobs")
       case 'q', 'Q':
           app.Stop()
       }
       return event
   })

   return app.SetRoot(layout, true).EnableMouse(true).Run()
}
