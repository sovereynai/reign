package api

import (
   "encoding/json"
   "io/ioutil"
   "net/http"
   "net/http/httptest"
   "os"
   "path/filepath"
   "testing"
)

func TestSubmitJob_Success(t *testing.T) {
   // Prepare a temporary image file
   tmpDir := t.TempDir()
   imgPath := filepath.Join(tmpDir, "img.txt")
   content := []byte("dummydata")
   if err := os.WriteFile(imgPath, content, 0644); err != nil {
       t.Fatalf("Failed to write temp image: %v", err)
   }

   expectedReq := JobRequest{Image: string(content), Model: "mymodel", Redundancy: 2}
   respBody := JobResponse{JobID: "1234", Hash: "abcd"}

   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       if r.URL.Path != "/jobs" {
           t.Errorf("Expected POST to /jobs, got %s", r.URL.Path)
       }
       // Read request
       body, err := ioutil.ReadAll(r.Body)
       if err != nil {
           t.Fatalf("Failed to read request body: %v", err)
       }
       var jr JobRequest
       if err := json.Unmarshal(body, &jr); err != nil {
           t.Fatalf("Invalid JSON: %v", err)
       }
       if jr != expectedReq {
           t.Errorf("Expected request %+v, got %+v", expectedReq, jr)
       }
       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(respBody)
   }))
   defer ts.Close()
   BaseURL = ts.URL

   res, err := SubmitJob(imgPath, "mymodel", 2)
   if err != nil {
       t.Fatalf("SubmitJob failed: %v", err)
   }
   if res.JobID != respBody.JobID || res.Hash != respBody.Hash {
       t.Errorf("Unexpected response: %+v", res)
   }
}

func TestSubmitJob_ErrorStatus(t *testing.T) {
   tmpDir := t.TempDir()
   imgPath := filepath.Join(tmpDir, "img.txt")
   os.WriteFile(imgPath, []byte("data"), 0644)
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       http.Error(w, "fail", http.StatusBadRequest)
   }))
   defer ts.Close()
   BaseURL = ts.URL
   if _, err := SubmitJob(imgPath, "m", 1); err == nil {
       t.Error("Expected error on bad status, got nil")
   }
}