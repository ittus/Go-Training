package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", showRegistrationPage)

		userRoutes.POST("/register", register)
	}
	router.GET("/article/view/:article_id", getArticle)

}
