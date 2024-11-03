package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/Molina-Q/svelte-wiki/backend/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.setReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.println("Starting Api")

	err := http.ListenAndServe("localhost:8000", r)
}