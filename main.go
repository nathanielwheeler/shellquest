package main

import (
  "embed"
  "fmt"
  "os"

  "github.com/nathanielwheeler/shellquest/server"
)

//go:embed public templates
var fs embed.FS

func main() {
  if err := server.Run(&fs); err != nil {
    fmt.Fprintf(os.Stdout, "shellquest: %s\n", err)
    os.Exit(1)
  }
}