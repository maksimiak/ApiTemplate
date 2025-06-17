package main

import (
	"example/hello/internal/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API service...")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	} else {
		log.Info("Server started successfully on localhost:8080")
	}

}
