package app

import (
	"net/http"

	"github.com/eggysetiawan/gorahm-starterkit/service"
	"github.com/gorilla/mux"
)


type UserHandlers struct {
	service service.UserService
}


func (uh UserHandlers) indexUser(w http.ResponseWriter, r *http.Request)  {
	// status := r.URL.Query().Get("status")
	
	users, err := uh.service.GetAllUser()

	if err != nil {
		writeResponse(w, r.Header.Get("Content-Type"), err.Code, err.AsMessage())
	} else {
		writeResponse(w, r.Header.Get("Content-Type"), http.StatusOK, users)
	}
}

func (uh UserHandlers)showUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	id := vars["user"]
	
	user, err := uh.service.GetUser(id)

	if err != nil {
		writeResponse(w, r.Header.Get("Content-Type"), err.Code, err.AsMessage())
	} else {
		writeResponse(w, r.Header.Get("Content-Type"), http.StatusOK, user)
	}
}