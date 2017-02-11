package main

import (
	"fmt"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Println(r.Method, r.RequestURI)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main Handler")
	fmt.Println(r.Method, r.RequestURI)
}

func main() {

	port := ":3000"

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(port, nil)
	fmt.Println("Go Server is now running on port", port)
}
