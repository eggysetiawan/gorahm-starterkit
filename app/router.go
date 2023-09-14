package app

import (
	"log"
	"net/http"

	"github.com/eggysetiawan/gorahm-starterkit/domain"
	"github.com/eggysetiawan/gorahm-starterkit/logger"
	"github.com/eggysetiawan/gorahm-starterkit/service"
	"github.com/gorilla/mux"
)

func Router() {
	logger.Info("Starting application...")

	router := mux.NewRouter()

	uh := UserHandlers{service.NewUserService(domain.NewUserRepositoryDb())}

	router.HandleFunc("/users", uh.indexUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{user}", uh.showUser).Methods(http.MethodGet)


	log.Fatal(http.ListenAndServe("localhost:3001", router))

}
