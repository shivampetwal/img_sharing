package api

import (
	"net/http"

	"code/idk/internal/api/handlers"
)

// RegisterRoutes wires up all your URL paths to handler functions.
func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// health‐check or home
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("anh"))
	})

	// your API endpoints
	mux.HandleFunc("/api/get-random-image", handlers.GetData)
	//mux.HandleFunc("/api/count", handlers.CountHandler)
	// ... add more routes here ...

	mux.HandleFunc("/api/upload", handlers.UploadImage)
	//mux.HandleFunc("/api/count", handlers.CountHandler)
	// ... add more routes here ...

	return mux
}
