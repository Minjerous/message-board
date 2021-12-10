package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	engine := gin.Default()

	//这里是jwt鉴权使用的登录接口  由于token输入太麻烦就实现一处就ok cookie 比较方便
	engine.POST("/jwtlogin", UserLogin)

	engine.POST("/register", register)
	engine.POST("/login", login)
	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.PUT("/password", changePassword)
	}

	postGroup := engine.Group("/post")
	{
		postGroup.POST("/addpost", auth, addPost)
		//postGroup.POST("/addpost", JWTAuthMiddleware(), addPost) //Jwt 测试项目
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
