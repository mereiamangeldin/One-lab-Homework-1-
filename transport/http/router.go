package http

func (s *Server) InitRoutes() {
	v1 := s.router.Group("api/v1")
	v1.Use(s.handler.UserIdentity)
	auth := s.router.Group("/auth")
	book := s.router.Group("/books")
	auth.POST("/sign-in", s.handler.SignIn)
	auth.POST("/sign-up", s.handler.SignUp)
	v1.PUT("/:id", s.handler.UpdateUser)
	v1.GET("/:id", s.handler.GetUserById)
	v1.PUT("/:id/change-password", s.handler.UpdatePassword)
	v1.DELETE("/:id", s.handler.DeleteUser)
	v1.GET("/:id/books", s.handler.GetUserBooks)
	v1.POST("/:id/books", s.handler.TakeBook)
	v1.DELETE("/:id/books/:book_id", s.handler.ReturnBook)

	book.GET("", s.handler.GetBooks)
	book.POST("", s.handler.CreateBook)
	book.GET("/:id", s.handler.GetBookById)
	book.PUT("/:id", s.handler.UpdateBook)
	book.DELETE("/:id", s.handler.DeleteBook)

}
