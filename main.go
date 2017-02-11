package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Println(r.Method, r.RequestURI)
}

func main() {

	port := ":3000"

	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
	fmt.Println("Go Server is now running on port", port)
}
