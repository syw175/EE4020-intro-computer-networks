package main

import (
  "fmt"
  "io/ioutils"
  "net/http"
  "net/url"
  "strings"
  "os"
)

// Func to return the contents of a file in the current directory
func listFile() []string {
  var filename []string
  // Return the files in the current directory
  entries, err := ioutil.ReadDir("./")
  if (err != nil) {
	os.panic()
  }

  fmt.Println(entries)
  for _, ent := range entries {
	filename = append(filename, ent.Name())
  }
  return filename
}

// Function to handle server requests
func customHandler() {
}

// Main function
func main() {
  fmt.Println("Launching server... ")
  // Open the server for requests at port 12001


}


