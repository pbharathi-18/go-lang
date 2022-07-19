package main

import (
	"net/http"
	"os"
)

//The w parameter is the structure we use to send responses to an HTTP request. It implements a Write() method which accepts a slice of bytes and writes the data to the connection as part of an HTTP response.
//the r parameter represents the HTTP request received from the client. It’s how we access the data sent by a client to the server.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}

//Getenv() returns an empty string and port is set to 3000 so that the server is made available at http://localhost:3000
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	//the http.NewServeMux() method is used to create an HTTP request multiplexer which is subsequently assigned to the mux variable
	mux := http.NewServeMux()
	//the http.NewServeMux() method is used to create an HTTP request multiplexer which is subsequently assigned to the mux variable.
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
