package cmd

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/http/httptest"
   "testing"

   "github.com/Leathal1/greycli/internal/api"
)


func TestJobsCmd_Text(t *testing.T) {
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       json.NewEncoder(w).Encode([]api.Job{{ID: "1", Model: "m", Status: "s"}})
   }))
   defer ts.Close()
   api.BaseURL = ts.URL
   out, err := executeTest("jobs")
   if err != nil {
       t.Fatalf("jobs command failed: %v", err)
   }
   if !bytes.Contains([]byte(out), []byte("🧾 Recent Jobs:")) {
       t.Errorf("Unexpected output: %s", out)
   }
}

func TestJobsCmd_JSON(t *testing.T) {
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       json.NewEncoder(w).Encode([]api.Job{{ID: "2", Model: "n", Status: "t"}})
   }))
   defer ts.Close()
   api.BaseURL = ts.URL
   out, err := executeTest("jobs", "--json")
   if err != nil {
       t.Fatalf("jobs --json failed: %v", err)
   }
   var data []api.Job
   if err := json.Unmarshal([]byte(out), &data); err != nil {
       t.Errorf("Invalid JSON output: %v", err)
   }
   if len(data) != 1 || data[0].ID != "2" {
       t.Errorf("Unexpected data: %v", data)
   }
}