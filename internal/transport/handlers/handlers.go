package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"onelab/internal/service"

	"github.com/gorilla/mux"
)

// func CreateUserHandler(s *service.Service) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var input struct {
// 			Name     string `json:"name"`
// 			Login    string `json:"username"`
// 			Password string `json:"password"`
// 		}
// 		err := json.NewDecoder(r.Body).Decode(&input)
// 		if err != nil {
// 			http.Error(w, "Invalid request body", http.StatusBadRequest)
// 			return
// 		}

// 		user := model.User{
// 			Name:     input.Name,
// 			Login:    input.Login,
// 			Password: input.Password,
// 		}

// 		err = s.User.Create(&user)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		resp := map[string]interface{}{
// 			"id":    user.ID,
// 			"name":  user.Name,
// 			"login": user.Login,
// 		}
// 		json.NewEncoder(w).Encode(resp)
// 	}
// }

func GetUserByIDHandler(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userIDStr := vars["id"]
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		user, err := s.User.Get(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUser)
	}
}
