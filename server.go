package main

import (
	"fmt"
	"net/http"
)

func handleHelloRoute(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc("/hello", handleHelloRoute)

	http.ListenAndServe(":5000", nil)
}
