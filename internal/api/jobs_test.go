package api

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/http/httptest"
   "testing"
)

func TestFetchJobs_Success(t *testing.T) {
   sample := []Job{{ID: "1", Model: "m1", Status: "ok"}}
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       if r.URL.Path != "/jobs" {
           t.Errorf("Expected request to /jobs, got %s", r.URL.Path)
       }
       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(sample)
   }))
   defer ts.Close()
   BaseURL = ts.URL
   jobs, err := FetchJobs()
   if err != nil {
       t.Fatalf("FetchJobs failed: %v", err)
   }
   if len(jobs) != len(sample) {
       t.Fatalf("Expected %d jobs, got %d", len(sample), len(jobs))
   }
   for i, j := range jobs {
       if j != sample[i] {
           t.Errorf("Job[%d]: expected %+v, got %+v", i, sample[i], j)
       }
   }
}

func TestFetchJobs_BadStatus(t *testing.T) {
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       w.WriteHeader(http.StatusBadGateway)
   }))
   defer ts.Close()
   BaseURL = ts.URL
   if _, err := FetchJobs(); err == nil {
       t.Error("Expected error on bad status code, got nil")
   }
}

func TestFetchJobs_InvalidJSON(t *testing.T) {
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       w.WriteHeader(http.StatusOK)
       fmt.Fprint(w, "not json")
   }))
   defer ts.Close()
   BaseURL = ts.URL
   if _, err := FetchJobs(); err == nil {
       t.Error("Expected JSON unmarshal error, got nil")
   }
}