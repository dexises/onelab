package transport

import (
	"net/http"

	"onelab/internal/service"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userManager *service.Service) http.Handler {
	router := httprouter.New()

	router.POST("/user", createUserHandler)
	router.GET("/user/:id", showUserHandler)

	return r
}
