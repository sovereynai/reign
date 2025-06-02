package cmd

import (
   "bytes"
   "io"
   "os"
)

// executeTest runs the root command with args, capturing stdout and stderr.
func executeTest(args ...string) (string, error) {
   oldOut, oldErr := os.Stdout, os.Stderr
   r, w, _ := os.Pipe()
   os.Stdout, os.Stderr = w, w
   defer func() { os.Stdout, os.Stderr = oldOut, oldErr }()

   rootCmd.SetArgs(args)
   err := rootCmd.Execute()

   w.Close()
   var buf bytes.Buffer
   io.Copy(&buf, r)
   return buf.String(), err
}