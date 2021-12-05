package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	engine := gin.Default()

	engine.POST("/login", login)
	engine.POST("/register", register)

	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.PUT("/password", changePassword)
	}

	postGroup := engine.Group("/post")
	{
		postGroup.Use(auth)
		postGroup.POST("/addpost", addPost)
		postGroup.DELETE("/deletePost", deletePost)
		postGroup.GET("/getpost", getPosts)
		postGroup.GET("/:post_id")
	}
	//评论楼中楼
	commentGroup := engine.Group("/comment")
	{
		commentGroup.Use(auth)
		commentGroup.POST("/")
		commentGroup.DELETE("/")
		commentGroup.GET("/")
	}

	engine.Run(":8090")
}
