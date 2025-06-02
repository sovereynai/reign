package cmd

import (
   "bytes"
   "encoding/json"
   "io/ioutil"
   "net/http"
   "net/http/httptest"
   "os"
   "path/filepath"
   "testing"

   "github.com/Leathal1/greycli/internal/api"
)


func TestSubmitCmd_Success(t *testing.T) {
   // Prepare temp image
   tmpDir := t.TempDir()
   imgPath := filepath.Join(tmpDir, "img.txt")
   data := []byte("imgdata")
   if err := os.WriteFile(imgPath, data, 0644); err != nil {
       t.Fatalf("Failed to write temp image: %v", err)
   }

   // Mock API
   expectedReq := api.JobRequest{Image: string(data), Model: "mod", Redundancy: 3}
   respData := api.JobResponse{JobID: "42", Hash: "hsh"}
   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       if r.URL.Path != "/jobs" {
           t.Errorf("Expected POST to /jobs, got %s", r.URL.Path)
       }
       body, _ := ioutil.ReadAll(r.Body)
       var jr api.JobRequest
       if err := json.Unmarshal(body, &jr); err != nil {
           t.Fatalf("Invalid JSON: %v", err)
       }
       if jr != expectedReq {
           t.Errorf("Expected request %+v, got %+v", expectedReq, jr)
       }
       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(respData)
   }))
   defer ts.Close()
   api.BaseURL = ts.URL

   // Run CLI
   out, err := executeTest(
       "submit",
       "--model", "mod",
       "--image", imgPath,
       "--redundancy", "3",
   )
   if err != nil {
       t.Fatalf("submit command failed: %v", err)
   }
   // Check output
   if !bytes.Contains([]byte(out), []byte("Job submitted successfully!")) {
       t.Errorf("Unexpected output: %s", out)
   }
   if !bytes.Contains([]byte(out), []byte("Job ID: 42")) {
       t.Errorf("Missing job ID, got: %s", out)
   }
}