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

func postSites(w http.ResponseWriter, r *http.Request) {
	conn := connectToServer()
	defer conn.Close()

	request := protocols.CreateSiteRequest{
		Command: "create-site",
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

	var response protocols.CreateSiteResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	fmt.Printf("%+v\n", response)
}

func getFiles(w http.ResponseWriter, r *http.Request) {
	conn := connectToServer()
	defer conn.Close()

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
