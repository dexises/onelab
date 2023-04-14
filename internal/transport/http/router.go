package http

func (s *Server) NewRouter() {
	v1 := s.App.Group("/api/v1")

	v1.POST("/login", s.handler.LoginUser)

	users := v1.Group("/users")
	users.POST("/", s.handler.CreateUser)
	users.GET("/:id", s.handler.GetUserByID)
	users.PUT("/:id", s.handler.UpdateUser)

	books := v1.Group("/books")
	books.POST("/", s.handler.CreateBook)
	books.GET("/:id", s.handler.GetBookByID)
	books.GET("/", s.handler.GetAllBook)

	bookReader := v1.Group("/library")
	//Получить текущих пользователей со списком книг который у них сейчас на руках
	bookReader.GET("/all", s.handler.List)
	//Список пользователей с количеством книг за последний месяц
	bookReader.GET("/month", s.handler.ListMonth)
	bookReader.POST("/give-book", s.handler.GiveBook)
	bookReader.POST("/return-book", s.handler.ReturnBook)
}
