package views

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var creditText *tview.TextView

func NewCreditsView() *tview.TextView {
	creditText = tview.NewTextView()
	creditText.SetText("Credits: 0").
		SetTextColor(tcell.ColorYellow).
		SetTitle("Credit Tracker").
		SetBorder(true)

	return creditText
}

func RefreshCredits() {
	resp, err := http.Get("http://localhost:8080/credits")
	if err != nil {
		creditText.SetText("Credits: ?")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		creditText.SetText("Credits: !")
		return
	}

	creditText.SetText(fmt.Sprintf("Credits: %s", body))
}
