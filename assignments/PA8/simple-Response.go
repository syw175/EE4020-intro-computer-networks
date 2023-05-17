package main

import "fmt"
import “net/http”

func main() {
  fmt.Println("Launching server...")
  http.ListenAndServe(":<your port#>", \
  http.FileServer(http.Dir(".")))
}
