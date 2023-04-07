package transport

import (
	"net/http"

	"onelab/internal/service"
	"onelab/internal/transport/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(svc *service.Service) http.Handler {
	r := mux.NewRouter()

	r.Use(logRequest)

	// r.HandleFunc("/users", handlers.CreateUserHandler(svc)).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUserByIDHandler(svc)).Methods("GET")
	return r
}
