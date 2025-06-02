package cmd

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/http/httptest"
   "testing"

   "github.com/Leathal1/greycli/internal/api"
)


func TestModelsCmd_Text(t *testing.T) {
   // Mock API
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       json.NewEncoder(w).Encode([]string{"a","b"})
   }))
   defer ts.Close()
   api.BaseURL = ts.URL

   out, err := executeTest("models")
   if err != nil {
       t.Fatalf("models command failed: %v", err)
   }
   if !bytes.Contains([]byte(out), []byte("🧠 Available models:")) {
       t.Errorf("Unexpected output: %s", out)
   }
}

func TestModelsCmd_JSON(t *testing.T) {
   // Mock API
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       json.NewEncoder(w).Encode([]string{"x","y"})
   }))
   defer ts.Close()
   api.BaseURL = ts.URL

   out, err := executeTest("models", "--json")
   if err != nil {
       t.Fatalf("models --json failed: %v", err)
   }
   var data []string
   if err := json.Unmarshal([]byte(out), &data); err != nil {
       t.Errorf("Invalid JSON output: %v", err)
   }
   if len(data) != 2 || data[0] != "x" || data[1] != "y" {
       t.Errorf("Unexpected JSON data: %v", data)
   }
}