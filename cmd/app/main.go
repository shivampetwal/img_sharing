package main

import (


	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 6009...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "main.goooooooooooooooo")
	})

	log.Fatal(http.ListenAndServe(":6009", nil))
}
