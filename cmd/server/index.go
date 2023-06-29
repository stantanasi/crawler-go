package main

import (
	"fmt"
	"net"
)

func main() {
	// configuration du serveur //
	_, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server started!")
	// configuration du serveur //
}
