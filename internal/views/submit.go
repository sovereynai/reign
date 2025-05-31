package views

import (
	"strconv"

	"github.com/Leathal1/greycli/internal/api"
	"github.com/rivo/tview"
)

func NewSubmitView() *tview.Form {
	form := tview.NewForm()

	imagePathField := tview.NewInputField().SetLabel("Image Path")
	modelDropdown := tview.NewDropDown().SetLabel("Model").SetOptions([]string{"mnml-alpha", "mnml-beta", "mnml-gamma"}, nil)
	redundancyField := tview.NewInputField().SetLabel("Redundancy").SetText("1")

	form.AddFormItem(imagePathField)
	form.AddFormItem(modelDropdown)
	form.AddFormItem(redundancyField)

	form.AddButton("Submit", func() {
		imagePath := imagePathField.GetText()
		_, modelStr := modelDropdown.GetCurrentOption()
		redundancy, _ := strconv.Atoi(redundancyField.GetText())

		resp, err := api.SubmitJob(imagePath, modelStr, redundancy)
		if err != nil {
			return
		}
		AddJobRow(resp.JobID, "Submitted", resp.Hash, modelStr)
		RefreshCredits()
	})

	form.AddButton("Cancel", func() {})
	form.SetTitle("Submit Modal").SetBorder(true)
	return form
}
