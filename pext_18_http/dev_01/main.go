package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	
	log.Println("Server start on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error on server start: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request ) {
	fmt.Fprintf(w, "<h1>Hello</h>")
}