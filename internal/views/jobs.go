package views

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var jobsTable *tview.Table

func NewJobsView() *tview.Table {
	jobsTable = tview.NewTable()
	jobsTable.SetBorders(true).
		SetTitle("Jobs").
		SetBorder(true)

	jobsTable.SetCell(0, 0, tview.NewTableCell("ID").SetTextColor(tcell.ColorWhite))
	jobsTable.SetCell(0, 1, tview.NewTableCell("Status").SetTextColor(tcell.ColorWhite))
	jobsTable.SetCell(0, 2, tview.NewTableCell("Proof").SetTextColor(tcell.ColorWhite))
	jobsTable.SetCell(0, 3, tview.NewTableCell("Model").SetTextColor(tcell.ColorWhite))

	return jobsTable
}

func AddJobRow(id, status, proof, model string) {
	row := jobsTable.GetRowCount()
	jobsTable.SetCell(row, 0, tview.NewTableCell(id).SetTextColor(tcell.ColorGreen))
	jobsTable.SetCell(row, 1, tview.NewTableCell(status).SetTextColor(tcell.ColorYellow))
	jobsTable.SetCell(row, 2, tview.NewTableCell(fmt.Sprintf("%.8s...", proof)).SetTextColor(tcell.ColorGray))
	jobsTable.SetCell(row, 3, tview.NewTableCell(model).SetTextColor(tcell.ColorWhite))
}
