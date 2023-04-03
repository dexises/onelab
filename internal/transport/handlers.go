package transport

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func createUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

func showUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}
