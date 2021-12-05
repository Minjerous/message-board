package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board-demo/model"
	"message-board-demo/service"
	"message-board-demo/tool"
	"time"
)

func getPosts(ctx *gin.Context) {
	posts, err := service.GetPosts()
	if err != nil {
		fmt.Println("get posts err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, posts)
}

func addPost(ctx *gin.Context) {
	txt := ctx.PostForm("txt")

	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	post := model.Post{
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}

	err := service.AddPost(post)
	if err != nil {
		fmt.Println("add post err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}
func deletePost(ctx *gin.Context) {
	//iUsername, _ := ctx.Get("username")
}
