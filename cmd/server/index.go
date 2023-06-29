package main

import (
	"fmt"
	"net"
	"time"
)

type Dir struct {
	id int64
	hostname string
	domain string
	lastseen time.Time
}

type File struct {
	id int64
	name string
	url string
	page string
	lastseen time.Time
}

type Database struct {
	dirs []Dir
	files []File
}

func main() {
	// configuration du serveur //
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
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
	}

}
