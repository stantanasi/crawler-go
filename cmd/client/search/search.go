package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stantanasi/crawler-go/pkg/protocols"
	"net"
	"net/http"
)

func main() {
	// Configure HTTP server
	r := mux.NewRouter()
	r.HandleFunc("/sites", getSites).Methods("GET")
	r.HandleFunc("/sites", postSites).Methods("POST")
	r.HandleFunc("/files", getFiles).Methods("GET")
	http.ListenAndServe(":8080", r)
}


func connectToServer() net.Conn {
	// configuration du serveur //
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return nil
	}

	return conn
}

func getSites(w http.ResponseWriter, r *http.Request) {
	conn := connectToServer()
	defer conn.Close()

	request := protocols.GetSiteRequest{
		Command: "get-site",
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, string(requestJSON))
}

func postSites(w http.ResponseWriter, r *http.Request) {
	conn := connectToServer()
	defer conn.Close()

	request := protocols.GetSiteRequest{
		Command: "create-site",
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, string(requestJSON))
}

func getFiles(w http.ResponseWriter, r *http.Request) {
	conn := connectToServer()
	defer conn.Close()
}
