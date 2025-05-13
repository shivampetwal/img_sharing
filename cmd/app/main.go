package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"code/idk/config"
	"code/idk/internal/api" // ← import your routes package
	"code/idk/internal/api/handlers"
	"code/idk/internal/db"

	_ "github.com/lib/pq"
)

func main() {
	// 1) load config & connect
	cfg := config.LoadConfig()
	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 2) give the handlers a DB handle
	handlers.SetDB(conn)

	// 3) mount all routes from internal/api/routes.go
	router := api.RegisterRoutes()

	// 4) start the server on PORT (default 6009)
	port := os.Getenv("PORT")
	if port == "" {
		port = "6009"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
