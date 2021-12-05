package api

import (
	"github.com/gin-gonic/gin"
	"message-board-demo/tool"
)

func auth(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		tool.RespErrorWithData(ctx, "请登陆后进行操作")
		ctx.Abort()
	}
	ctx.Set("username", username)
	ctx.Next()
}
