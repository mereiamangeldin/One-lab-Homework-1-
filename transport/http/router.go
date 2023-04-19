package http

func (s *Server) InitRoutes() {
	v1 := s.router.Group("api/v1")
	v1.Use(s.handler.UserIdentity)
	auth := s.router.Group("/auth")
	auth.POST("/sign-in", s.handler.SignIn)
	auth.POST("/sign-up", s.handler.SignUp)
	v1.PUT("/:id", s.handler.UpdateUser)
	v1.GET("/:id", s.handler.GetUserById)
	v1.PUT("/:id/change-password", s.handler.UpdatePassword)
	v1.DELETE("/:id", s.handler.DeleteUser)
	v1.GET("/:id/user-books", s.handler.GetUserBooks)
	v1.GET("/:id/balance", s.handler.GetBalance)
	//v1.DELETE("/:id/transactions/:book_id", s.handler.ReturnBook)
	v1.GET("/:id/transactions"+
		"+/:book_id", s.handler.ReturnBook)

	v1.GET("/:id/books", s.handler.GetBooks)
	v1.GET("/:id/books/:id", s.handler.GetBookById)
	v1.POST("/:id/books/:book_id/rent", s.handler.RentBook)
	v1.POST("/:id/books/:book_id/purchase", s.handler.BuyBook)

}
