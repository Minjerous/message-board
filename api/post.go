package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board-demo/model"
	"message-board-demo/service"
	"message-board-demo/tool"
	"strconv"
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
	id := ctx.PostForm("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		tool.RespErrorWithData(ctx, "删除失败")
		fmt.Println("00")
		fmt.Println(err)
	}
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	post := model.Post{
		Id:       ID,
		Username: username,
	}
	err = service.DeletePost(post)
	if err != nil {
		tool.RespErrorWithData(ctx, "删除失败")
		fmt.Println(err)
		fmt.Println("000")
		return
	}
	tool.RespSuccessfulWithData(ctx, "删除成功")
}
