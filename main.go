package main

import (
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	http.ListenAndServe(":"+port, nil)
}
