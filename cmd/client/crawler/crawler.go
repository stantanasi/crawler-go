package main

import (
	"fmt"
	"net"
)

func main() {
	// configuration du serveur //
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer conn.Close()
	// configuration du serveur //

}
