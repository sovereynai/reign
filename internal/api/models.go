package api

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "net/http"
)

// BaseURL is the HTTP base address for the API; override in tests as needed.
var BaseURL = "http://localhost:8080"

func FetchModels() ([]string, error) {
	resp, err := http.Get(BaseURL + "/models")
	if err != nil {
		return nil, fmt.Errorf("get models: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	var models []string
	if err := json.Unmarshal(body, &models); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return models, nil
}
