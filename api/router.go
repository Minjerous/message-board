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
		//游客状态下
		postGroup.POST("/addpost", auth, addPost)
		postGroup.DELETE("/deletePost", auth, deletePost)
		postGroup.GET("/getpost", getPosts)
	}
	//评论楼中楼
	commentGroup := engine.Group("/comment")
	{
		commentGroup.Use(auth)
		commentGroup.POST("/add", addComment)
		postGroup.POST("/anonymous", addAnonymousComment)
		commentGroup.DELETE("/delete", deleteComment)
		commentGroup.GET("/get", getComment)
		commentGroup.POST("/addcomment", AtCommentaddComment)
	}

	engine.Run(":8090")
}
