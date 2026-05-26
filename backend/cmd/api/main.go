package main

import (
	"log"
	"net/http"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/routes"
)

func main() {
	router := routes.NewRouter()

	log.Println("API running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
