package main

import (
	"encoding/json"
	"fmt"
	"github.com/stantanasi/crawler-go/pkg/protocols"
	"net"
	"time"
)

type Dir struct {
	id       int64
	hostname string
	domain   string
	lastseen time.Time
}

type File struct {
	id       int64
	name     string
	url      string
	page     string
	lastseen time.Time
}

type Database struct {
	dirs  []Dir
	files []File
}

func main() {
	// configuration du serveur //
	ln, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server started!")

	// boucle continue pour accepter les connexions //
	for {
		conn, err := ln.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		response := make(chan any)

		go handleRequest(conn, response)

		responseJSON, err := json.Marshal(<-response)
		if err != nil {
			fmt.Println("Error creating response:", err)
			return
		}

		_, err = conn.Write([]byte(string(responseJSON)))
		if err != nil {
			fmt.Println("Error sending response:", err)
			return
		}
	}
}

func handleRequest(conn net.Conn, response chan<- any) {
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	var request protocols.GenericRequest
	err = json.Unmarshal(buffer[:n], &request)
	if err != nil {
		fmt.Println("Error parsing request:", err)
		return
	}

	response <- protocols.GenericResponse{
		Command: request.Command,
		Status: "nok",
	}
}
