package api

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/http/httptest"
   "testing"
)

func TestFetchCredits_Success(t *testing.T) {
   sample := struct{ Credits int `json:"credits"`} {Credits: 42}
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       if r.URL.Path != "/credits" {
           t.Errorf("Expected request to /credits, got %s", r.URL.Path)
       }
       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(sample)
   }))
   defer ts.Close()
   BaseURL = ts.URL
   credits, err := FetchCredits()
   if err != nil {
       t.Fatalf("FetchCredits failed: %v", err)
   }
   if credits != sample.Credits {
       t.Errorf("Expected %d credits, got %d", sample.Credits, credits)
   }
}

func TestFetchCredits_BadStatus(t *testing.T) {
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       w.WriteHeader(http.StatusNotFound)
   }))
   defer ts.Close()
   BaseURL = ts.URL
   if _, err := FetchCredits(); err == nil {
       t.Error("Expected error on bad status code, got nil")
   }
}

func TestFetchCredits_InvalidJSON(t *testing.T) {
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       w.WriteHeader(http.StatusOK)
       fmt.Fprint(w, "bad json")
   }))
   defer ts.Close()
   BaseURL = ts.URL
   if _, err := FetchCredits(); err == nil {
       t.Error("Expected JSON unmarshal error, got nil")
   }
}