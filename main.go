package main

import (
   "github.com/Leathal1/greycli/cmd"
)

// version information set via ldflags
var (
   version = "dev"
   commit  string
   date    string
   builtBy string
)

func main() {
	cmd.Execute()
}
