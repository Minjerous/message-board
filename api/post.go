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

//查看文章
func getPosts(ctx *gin.Context) {
	posts, err := service.GetPosts()
	if err != nil {
		fmt.Println("get posts err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespAllPosts(ctx, posts)
}

//增加文章
func addPost(ctx *gin.Context) {
	txt := ctx.PostForm("txt")
	//判断是否含有敏感词
	if tool.CheckIfSensitive(txt) {
		tool.RespErrorWithData(ctx, "审核未通过，请注意你的言词")
		return
	}
	//评论不能为空
	if txt == "" {
		tool.RespErrorWithData(ctx, "不可以发表空评论")
		return
	}
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	post := model.Post{
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	//添加评论
	err := service.AddPost(post)
	if err != nil {
		fmt.Println("add post err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

//删除文章
func deletePost(ctx *gin.Context) {
	id := ctx.PostForm("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		tool.RespErrorWithData(ctx, "删除失败")
		fmt.Println(err)
		return
	}
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	post := model.Post{
		Id:         ID,
		Username:   username,
		UpdateTime: time.Now(),
	}
	flag, err := service.IsUsernameMachIdByPost(username, id)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Print(err)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "没有权限")
		return
	}
	err = service.DeletePost(post)
	if err != nil {
		tool.RespErrorWithData(ctx, "删除失败")
		fmt.Println(err)
		return
	} else {
		tool.RespSuccessfulWithData(ctx, "删除成功")
	}
}
