package main

import (
  "io"
  "os"
  "fmt"
)

func main() {
  filePath := os.Args[1]
  file, err := os.Open(filePath)
  if err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }
  io.Copy(os.Stdout, file)
}
