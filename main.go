package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	requestMSG := "GET / HTTP/1.1\r\n"
	requestMSG += "User-Agent: NTRIP Agent\r\n"
	requestMSG += "Accept: */*\r\n"
	requestMSG += "Connection: close\r\n"
	requestMSG += "\r\n"

	fmt.Println("string:", requestMSG)

	conn, err := net.Dial("tcp", "Your Address")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(requestMSG))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	fmt.Println("Request sent successfully.")

	// Read response
	response, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response received:", string(response))
}
