package main

import (
  "fmt"
  "net/http"
)

const port = "5010"

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, Booker")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":" + port, nil)
}
