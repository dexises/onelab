package http

func (s *Server) NewRouter() {
	// r.HandleFunc("/users", handlers.CreateUserHandler(svc)).Methods("POST")
	// r.HandleFunc("/users/{id}", handlers.GetUserByIDHandler(svc)).Methods("GET")

	s.App.POST("/users", s.handler.CreateUser)
	s.App.GET("/users/:id", s.handler.GetUserByID)
}
