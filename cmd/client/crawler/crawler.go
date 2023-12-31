package main

import (
	"encoding/json"
	"fmt"
	"github.com/stantanasi/crawler-go/pkg/protocols"
	"net"
)

func main() {
	// Configure TCP server
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer conn.Close()
}

func getSite(conn net.Conn) {
	request := protocols.GetSiteRequest{
		Command: "get-site",
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, string(requestJSON))

	// Réception de la réponse
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var response protocols.GetSiteResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	fmt.Printf("%+v\n", response)
}

func updateSite(conn net.Conn) {
	request := protocols.UpdateSiteRequest{
		Command: "update-site",
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, string(requestJSON))

	// Réception de la réponse
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var response protocols.UpdateSiteResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	fmt.Printf("%+v\n", response)
}

func getFile(conn net.Conn) {
	request := protocols.GetFileRequest{
		Command: "get-file",
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, string(requestJSON))

	// Réception de la réponse
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var response protocols.GetFileResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	fmt.Printf("%+v\n", response)
}

func createFile(conn net.Conn) {
	request := protocols.CreateFileRequest{
		Command: "create-file",
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, string(requestJSON))

	// Réception de la réponse
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var response protocols.CreateFileResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	fmt.Printf("%+v\n", response)
}

func updateFile(conn net.Conn) {
	request := protocols.UpdateFileRequest{
		Command: "update-file",
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, string(requestJSON))

	// Réception de la réponse
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var response protocols.UpdateFileResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	fmt.Printf("%+v\n", response)
}
