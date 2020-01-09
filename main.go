package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, Booker")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":5010", nil)
}
