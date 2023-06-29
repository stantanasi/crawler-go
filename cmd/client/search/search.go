package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

func main() {
	// Configure HTTP server
	r := mux.NewRouter()
	http.ListenAndServe(":8080", r)

	// configuration du serveur //
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer conn.Close()
	// configuration du serveur //
}
