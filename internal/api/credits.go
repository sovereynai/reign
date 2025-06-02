package api

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "net/http"
)


func FetchCredits() (int, error) {
	resp, err := http.Get(BaseURL + "/credits")
	if err != nil {
		return 0, fmt.Errorf("failed to fetch credits: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response: %w", err)
	}

	var result struct {
		Credits int `json:"credits"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("failed to parse credits: %w", err)
	}

	return result.Credits, nil
}
