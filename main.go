package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	} else if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	} else {
		fmt.Fprintf(w, "Hello World!")
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	} else if r.Method != "POST" {
		http.Redirect(w, r, "/form.html", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Post from website!")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s \n", name)
	fmt.Fprintf(w, "Address = %s \n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Server is running on port 8080")
	}
}