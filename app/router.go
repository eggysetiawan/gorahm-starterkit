package app

import (
	"github.com/eggysetiawan/go-email-blast/domain"
	"github.com/eggysetiawan/go-email-blast/logger"
	"github.com/eggysetiawan/go-email-blast/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() {
	logger.Info("Starting application...")

	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe("localhost:3001", router))

}
