package main

import (
	"fmt"
	"net/http"

	"github.com/ShallowMelon/GoAPI2/tools"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GoAPI Services...")

	fmt.Println("Welcome to GoAPI")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
