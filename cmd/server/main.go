package main

import (
	"log"
	"net/http"

	"github.com/prodownlaoder/Direct/router"
	"github.com/rs/cors"
)

func main() {
	r := router.SetupRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	log.Println("🚀 Server running at http://localhost:8082")
	if err := http.ListenAndServe(":8082", corsHandler); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
