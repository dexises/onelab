package http

func (s *Server) NewRouter() {
	s.App.POST("/", s.handler.HelloWorld)

	// s.App.POST("/users", s.handler.CreateUser)
	// s.App.GET("/users/:id", s.handler.GetUserByID)
}
