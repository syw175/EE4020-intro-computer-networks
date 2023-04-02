// Listen at 12001 until there's an upload request
// Read from the socket first the file size (given in a single line)
// Read from the socket one line at a time
// Prepend the line count to each line and store new line into a new file called: whatever.txt
// Repeat (3) and (4) until all lines in the file are processed
// Send a message to client the original and new file sizes
// Close the connection and terminate the program

package main

import (
	"fmt"
	"net"
)

// Function to check for errors, if not null, panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Launch the server on local ip at port 12001
	fmt.Println("Launching server... at port 12001")
	ln, _ := net.Listen("tcp", ":12001")
	conn, _ := ln.Accept()
	ln.Close()
	conn.Close()

}
