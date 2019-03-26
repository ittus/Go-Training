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
	articleRoutes := router.Group("/article")
	{
		// route from Part 1 of the tutorial
		articleRoutes.GET("/view/:article_id", getArticle)

		articleRoutes.GET("/create", showArticleCreationPage)

		articleRoutes.POST("/create", createArticle)
	}
}
