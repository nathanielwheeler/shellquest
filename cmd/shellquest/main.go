package main

import (
  "fmt"
  "os"

  "github.com/nathanielwheeler/shellquest/shell"
)

func main() {
  if err := shell.Run(); err != nil {
    fmt.Fprintf(os.Stdout, "shellquest: %s\n", err)
    os.Exit(1)
  }
}