// Objectives
// Prompt the user for upload file name
// Sends first the file size (# in size line)
// Send the file contents
// Receive a message back from server
// Print the server response
// Close the connection and terminate the program

package main

import (
	// "bufio"
	"net"
	// "os"
	// "fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Verify that this is the command to connect to pollys server
	conn, errc := net.Dial("tcp", "127.0.0.1:12000")
	check(errc)
	defer conn.Close()

}
