package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board-demo/dao"
	"message-board-demo/model"
	"message-board-demo/service"
	"message-board-demo/tool"
	"strconv"
	"time"
)

func addComment(ctx *gin.Context) {
	//获取文章ID
	PostId := ctx.PostForm("post_id")
	postId, _ := strconv.Atoi(PostId)
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
	comment := model.Comment{
		PostID:     postId,
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	err := service.AddNormalComment(comment)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "评论成功")
}

func getComment(ctx *gin.Context) {
	comments, err := service.GetComment()
	if err != nil {
		fmt.Println("get comment  err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, comments)
}

func deleteComment(ctx *gin.Context) {
	id := ctx.PostForm("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		tool.RespErrorWithData(ctx, "删除失败")
		fmt.Println(err)
		return
	}
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	comment := model.Comment{
		Id:       ID,
		Username: username,
	}
	flag, err := service.IsUsernameMachIdByComment(username, id)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Print(err)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "没有权限")
		return
	}
	err = service.DeleteComment(comment)
	if err != nil {
		tool.RespErrorWithData(ctx, "删除失败")
		fmt.Println(err)
		return
	} else {
		dao.MiuCommentNum(comment)
		tool.RespSuccessfulWithData(ctx, "删除成功")
	}
}

//丐版匿名评论  匿名评论可以直接在addComment 中获取一个评论状态但感觉测试麻烦 就早造了一个

func GetAnonymousComment(ctx *gin.Context) {
	PostId := ctx.PostForm("post_id")
	postId, _ := strconv.Atoi(PostId)
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
	comment := model.Comment{
		PostID:     postId,
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	err := service.AddAnonymousComment(comment)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "评论成功")
}
