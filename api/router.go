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

		postGroup.POST("/addpost", auth, addPost)
		postGroup.DELETE("/deletePost", auth, deletePost)
		//游客状态下可以获取文章信息
		postGroup.GET("/getpost", getPosts)
	}
	//评论楼中楼
	commentGroup := engine.Group("/comment")
	{
		commentGroup.Use(auth)
		//在文章下面进行评论
		commentGroup.POST("/add", addComment)
		//在文章下面进行匿名评论
		commentGroup.POST("/anonymous", addAnonymousComment)
		//在删除评论或回复
		commentGroup.DELETE("/delete", deleteComment)
		//获取同一文章下的所有评论以及回复
		commentGroup.GET("/getAll", getComment)
		//回复评论
		commentGroup.POST("/addcomment", atCommentAddComment)
		//获取单个评论的所有回复
		commentGroup.GET("/get", getOneCommentAllResp)
	}

	engine.Run(":8090")
}
