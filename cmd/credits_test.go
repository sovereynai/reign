package cmd

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/http/httptest"
   "testing"

   "github.com/Leathal1/greycli/internal/api"
)


func TestCreditsCmd_Text(t *testing.T) {
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       json.NewEncoder(w).Encode(map[string]int{"credits": 99})
   }))
   defer ts.Close()
   api.BaseURL = ts.URL
   out, err := executeTest("credits")
   if err != nil {
       t.Fatalf("credits command failed: %v", err)
   }
   if !bytes.Contains([]byte(out), []byte("💰 Current Credits: 99")) {
       t.Errorf("Unexpected output: %s", out)
   }
}

func TestCreditsCmd_JSON(t *testing.T) {
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       json.NewEncoder(w).Encode(map[string]int{"credits": 7})
   }))
   defer ts.Close()
   api.BaseURL = ts.URL
   out, err := executeTest("credits", "--json")
   if err != nil {
       t.Fatalf("credits --json failed: %v", err)
   }
   var data map[string]int
   if err := json.Unmarshal([]byte(out), &data); err != nil {
       t.Errorf("Invalid JSON output: %v", err)
   }
   if data["credits"] != 7 {
       t.Errorf("Unexpected data: %v", data)
   }
}