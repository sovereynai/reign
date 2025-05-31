package views

import (
	"github.com/Leathal1/greycli/internal/api"
	"github.com/rivo/tview"
)

func NewModelsView() *tview.List {
	modelList := tview.NewList()
	modelList.SetTitle("Model Selector").
		SetBorder(true)

	models, err := api.FetchModels()
	if err != nil {
		modelList.AddItem("Error loading models", "", 0, nil)
		return modelList
	}

	for _, model := range models {
		modelList.AddItem(model, "", 0, nil)
	}

	return modelList
}
