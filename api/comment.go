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
	err := service.AddComment(comment)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "评论成功")
}
