package main

import (
	"code/idk/config"
	"code/idk/internal/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 6009...")

	cfg := config.LoadConfig()


/******** DB CONNECT + PING *********/
_, err := db.Connect(cfg)
if err != nil {
	log.Fatalf("Failed to connect to database: %v", err)
}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "main.goooooooooooooooo")
	})

	log.Fatal(http.ListenAndServe(":6009", nil))
}
