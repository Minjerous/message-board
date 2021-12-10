package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board-demo/model"
	"message-board-demo/service"
	"message-board-demo/tool"
	"net/http"
)

//登录
func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	//var user model.User
	//err := ctx.ShouldBind(&user)
	flag, err := service.IsPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password correct err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithData(ctx, "密码错误")
		return
	}

	ctx.SetCookie("username", username, 60, "/", "", false, false)
	tool.RespSuccessful(ctx)
}

func UserLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	flag, err := service.IsPasswordCorrect(username, password)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("UserLogin is", err)
		return
	}
	// 校验用户名和密码是否正确
	if flag {
		// 生成Token
		tokenString, _ := tool.GenToken(username)
		tool.RespErrorWithData(ctx, tokenString)
		return
	}
	tool.RespErrorWithData(ctx, "鉴权失败")
	return
}

//注册
func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if len(username) < 3 {
		tool.RespErrorWithData(ctx, "用户名至少两位")
		return
	}

	if len(password) < 8 {
		tool.RespErrorWithData(ctx, "密码必须大于8位")
		return
	}

	//加盐加密
	passWord := tool.HashWithSalted(password)
	user := model.User{
		Username: username,
		Password: passWord,
	}

	//判断是否用户名已经被注册
	flag, err := service.IsRepeatUsername(user.Username)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}

	if flag {
		tool.RespErrorWithData(ctx, "用户名重复")
		return
	}

	err = service.Register(user)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	} else {
		tool.RespSuccessful(ctx)
	}
}

//修改密码
func changePassword(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	FirstNewPassword := ctx.PostForm("newpasswordOne")
	SecondNewPassword := ctx.PostForm("newpasswordTwo")

	user := model.User{
		Username: username,
		Password: tool.HashWithSalted(FirstNewPassword),
	}

	flag, err := service.IsRepeatUsername(username)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}

	if flag == false {
		tool.RespErrorWithData(ctx, "用户不存在")
		return
	}
	if flag {
		flag, err := service.IsPasswordCorrect(username, password)
		if err != nil {
			fmt.Println(err)
			tool.RespInternalError(ctx)
			return
		}

		if flag {
			if FirstNewPassword == SecondNewPassword {
				err := service.Password(user)
				if err == nil {
					tool.RespSuccessfulWithData(ctx, "修改成功")
					return
				} else {
					tool.RespErrorWithData(ctx, "修改失败")
					fmt.Println("err by change password is", err)
					return
				}
				return
			} else {
				tool.RespErrorWithData(ctx, "两次密码输入错误")
				return
			}
		}
		return
	}
}
func authHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	if user.Username == "q1mi" && user.Password == "q1mi123" {
		// 生成Token
		tokenString, _ := tool.GenToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}
