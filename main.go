package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func loadPage(title string) string {
	filename := title + ".html"
	body, _ := ioutil.ReadFile(filename)
	return string(body)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Println(r.Method, r.RequestURI)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, loadPage("index"))
	fmt.Println(r.Method, r.RequestURI)
}

func main() {

	port := ":3000"

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(port, nil)
	fmt.Println("Go Server is now running on port", port)
}
