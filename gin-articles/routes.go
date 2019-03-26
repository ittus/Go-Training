package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", showRegistrationPage)

		userRoutes.POST("/register", register)
		userRoutes.GET("/login", showLoginPage)
		userRoutes.POST("/login", performLogin)
		userRoutes.GET("/logout", logout)
	}
	router.GET("/article/view/:article_id", getArticle)

}
