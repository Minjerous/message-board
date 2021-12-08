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

//发表评论
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
	pidName, err := dao.SelectPostUserByPostIdByPost(postId)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}
	comment := model.Comment{
		PostID:     postId,
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
		PidName:    pidName,
	}
	err = service.AddNormalCommentAtPost(comment)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "评论成功")
}

//查看post的评论
func getComment(ctx *gin.Context) {
	PostId := ctx.PostForm("post_id")
	postId, _ := strconv.Atoi(PostId)
	comments, err := service.GetComment(postId)
	if err != nil {
		fmt.Println("get comment  err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, comments)
}

//删除评论
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
		Id:         ID,
		Username:   username,
		UpdateTime: time.Now(),
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
		c, err := dao.SelectPostIdByIdByComment(comment.Id)
		if err != nil {
			tool.RespInternalError(ctx)
			fmt.Println("SelectPostIdByIdByComment is", err)
			return
		}
		dao.MiuCommentNum(c)
		tool.RespSuccessfulWithData(ctx, "删除成功")
	}
}

//匿名评论
func addAnonymousComment(ctx *gin.Context) {
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

//回复评论
func AtCommentaddComment(ctx *gin.Context) {
	//获取文章ID
	id := ctx.PostForm("id")
	Id, _ := strconv.Atoi(id)
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
	pidName, err := dao.SelectCommentUserByIdByComment(Id)
	c, err := dao.SelectPostIdByIdByComment(Id)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}
	comment := model.Comment{
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
		PidName:    pidName,
		PostID:     c.PostID,
		Id:         Id,
	}
	err = service.AddNormalCommentAtComment(comment)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "评论成功")
}
