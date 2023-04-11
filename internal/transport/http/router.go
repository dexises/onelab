package http

func (s *Server) NewRouter() {
	s.App.POST("/users", s.handler.CreateUser)
	s.App.GET("/users/:id", s.handler.GetUserByID)
	s.App.PUT("/users/:id", s.handler.UpdateUser)
}
