package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board-demo/model"
	"message-board-demo/service"
	"message-board-demo/tool"
)

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
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

	ctx.SetCookie("username", username, 600000, "/", "", false, false)
	tool.RespSuccessful(ctx)
}

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

	passWord := tool.HashWithSalted(password)
	user := model.User{
		Username: username,
		Password: passWord,
	}

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
