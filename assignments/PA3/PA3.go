// Objectives
// Sends first the file size (# in size line)
// Send the file contents
// Receive a message back from server
// Print the server response
// Close the connection and terminate the program

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

	// Prompt the user for upload file name
	uploadFileName := ""
	fmt.Print("Please enter the upload file name: ")
	fmt.Scanf("%s", &uploadFileName)

	// Open the file
	file, errc := os.Open(uploadFileName)
	check(errc)
	defer file.Close()
	// Calculate its file size
	stat, errc := file.Stat()
	check(errc)
	sizeBytes := stat.Size()

	// Create a buffered writer to send the file size
	writer := bufio.NewWriter(conn)
	// Send the file size
	_, errw := writer.WriteString(fmt.Sprintf("%d", sizeBytes))
	check(errw)
	writer.Flush()

	// Read the file contents and send it to the server at once
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, errw := writer.WriteString(scanner.Text())
		check(errw)
	}
	writer.Flush()

	// Read the server response
	scanner = bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Close the connection and terminate the program
	conn.Close()
}
