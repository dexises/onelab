package transport

import (
	"net/http"

	"onelab/internal/service"

	"github.com/gorilla/mux"
)

func NewRouter(svc *service.Service) http.Handler {
	r := mux.NewRouter()

	r.Use(logRequest)

	r.HandleFunc("/users", createUserHandler(svc)).Methods("POST")
	r.HandleFunc("/users/{id}", getUserByIDHandler(svc)).Methods("GET")
	return r
}
