package api

import (
	"github.com/gin-gonic/gin"
	"message-board-demo/tool"
	"net/http"
	"strings"
)

//cookie 验证登入
func auth(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		tool.RespErrorWithData(ctx, "请登陆后进行操作")
		ctx.Abort()
	}
	ctx.Set("username", username)
	ctx.Next()
}

//jwt
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		mc, err := tool.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}

		c.Set("username", mc.Username)
		c.Next()
	}
}
