package http

func (s *Server) NewRouter() {
	s.App.GET("/transactions", s.handler.AllTransactions)
	s.App.POST("/transactions", s.handler.TransactionsCreate)
}
