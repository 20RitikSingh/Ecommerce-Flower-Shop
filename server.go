package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Define the directory to serve files from
	dir := "./"

	// Define the entry point file
	entryPoint := "index.html"

	// Create a file server handler to serve static files
	fs := http.FileServer(http.Dir(dir))

	// Define a custom handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the requested file path
		requestedPath := r.URL.Path

		// If the requested path is for the root, serve the entry point file
		if requestedPath == "/" {
			http.ServeFile(w, r, filepath.Join(dir, entryPoint))
			return
		}

		// Otherwise, serve the requested file using the file server handler
		fs.ServeHTTP(w, r)
	})

	http.Handle("/metrics", promhttp.Handler())

	// Start the server
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
