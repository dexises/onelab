package http

import echoSwagger "github.com/swaggo/echo-swagger"
import _ "onelab/docs"

func (s *Server) NewRouter() {
	s.App.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	v1 := s.App.Group("/api/v1")
	onlyForAuthUsers := s.handler.JWTAuth.ValidateAuth

	users := v1.Group("/users")
	users.POST("/create", s.handler.CreateUser)
	users.POST("/login", s.handler.LoginUser)
	users.GET("/:id", s.handler.GetUserByID, onlyForAuthUsers)
	users.PUT("/:id", s.handler.UpdateUser, onlyForAuthUsers)

	books := v1.Group("/books")
	books.POST("/", s.handler.CreateBook, onlyForAuthUsers)
	books.GET("/:id", s.handler.GetBookByID)
	books.GET("/", s.handler.GetAllBook)

	bookReader := v1.Group("/library")
	//Получить текущих пользователей со списком книг который у них сейчас на руках
	bookReader.GET("/all", s.handler.List, onlyForAuthUsers)
	//Список пользователей с количеством книг за последний месяц
	bookReader.GET("/month", s.handler.ListMonth, onlyForAuthUsers)
	bookReader.POST("/give-book", s.handler.GiveBook, onlyForAuthUsers)
	bookReader.POST("/return-book", s.handler.ReturnBook, onlyForAuthUsers)
}
