package main

import (
  "fmt"
  "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello, world!")
}

func main() {
 
  fmt.Println("Launching server...")
  hh := http.HandlerFunc(helloHandler)
  http.Handle("/hello", hh)
  fs := http.FileServer(http.Dir("."))
  http.Handle("/", http.StripPrefix("/", fs))
  http.ListenAndServeTLS(":12001", "server.cer", "server.key", nil)
}

