package main

import "fmt"

func main() {
  fmt.Printf("Who's there?\n")
  text := ""
  fmt.Scanf("%s", &text)
  fmt.Printf("Hello, %s", text)
}

