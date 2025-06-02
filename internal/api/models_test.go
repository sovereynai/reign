package api

import (
   "encoding/json"
   "net/http"
   "net/http/httptest"
   "testing"
)

func TestFetchModels_Success(t *testing.T) {
   // Mock server returning two models
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       if r.URL.Path != "/models" {
           t.Errorf("Expected request to /models, got %s", r.URL.Path)
       }
       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode([]string{"foo", "bar"})
   }))
   defer ts.Close()
   BaseURL = ts.URL
   models, err := FetchModels()
   if err != nil {
       t.Fatalf("FetchModels failed: %v", err)
   }
   expected := []string{"foo", "bar"}
   if len(models) != len(expected) {
       t.Fatalf("Expected %d models, got %d", len(expected), len(models))
   }
   for i, m := range models {
       if m != expected[i] {
           t.Errorf("Model[%d]: expected %s, got %s", i, expected[i], m)
       }
   }
}

func TestFetchModels_BadStatus(t *testing.T) {
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       http.Error(w, "err", http.StatusInternalServerError)
   }))
   defer ts.Close()
   BaseURL = ts.URL
   if _, err := FetchModels(); err == nil {
       t.Error("Expected error on bad status code, got nil")
   }
}