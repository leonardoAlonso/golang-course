package main

import (
  "net/http"
  "fmt"
  "os"
)

func main() {
  resp, err := http.Get("https://www.google.com")
  if err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }
  fmt.Println(resp)  
}
