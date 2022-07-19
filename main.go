// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	var grades int = 45
// 	var message string = "hello world"

// 	fmt.Printf("variable grades= %v is of type %v \n", grades, reflect.TypeOf(grades))
// 	fmt.Printf("variable message=%v is of type '%v' \n", message, reflect.TypeOf(message))
// }
package main

import (
	"fmt"
	"log"
	"net/http"
)

// net/http package is a struct that contains many fields that makes a complete Request.
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " hello world !!! ")

}

//The Handler request net/http package wraps serveHTTP Function that takes Response and Pointer to Request.
func handleRequest() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}

func main() {
	handleRequest()
}

//we attach the server HTTP method to it and  the type webServer also starts behaving like a Handler and we can pass it to the ListenAndServe Function that takes an address and the Handler to Handle.
