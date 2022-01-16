package api

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// changePassword 改密码
func changePassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	iUsername, _ := ctx.Get("username")
	l1 := len([]rune(newPassword))
	if l1 <= 16 && l1 >= 6 { //强制规定密码小于16位并大于6位
		username := iUsername.(string)

		//检验旧密码是否正确
		flag, err := service.IsPasswordCorrect(username, oldPassword)
		if err != nil {
			fmt.Println("judge password correct err: ", err)
			tool.RespInternalError(ctx)
			return
		}

		if !flag {
			tool.RespErrorWithDate(ctx, "旧密码输入错误")
			return
		}

		//修改新密码
		err = service.ChangePassword(username, newPassword)
		if err != nil {
			fmt.Println("change password err: ", err)
			tool.RespInternalError(ctx)
			return
		}

		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "密码请在6位到16位之内")
		return
	}
}

// login 登录
func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	flag, err := service.IsPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password correct err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	//密保判断 能用,不过显示有点小问题
	if !flag {
		tool.RespErrorWithDate(ctx, "密码错误")
		return
	}

	//jwt
	c := model.MyClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "YuanXinHao",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessfulWithDate(ctx, s)
}

// register 注册
func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	question := ctx.PostForm("question")
	answer := ctx.PostForm("answer")
	//输入信息不能为空
	if username != "" && password != "" && question != "" && answer != "" {
		l1 := len([]rune(username))
		l2 := len([]rune(password))
		if l1 <= 8 && l1 >= 1 { //强制规定用户名长度小于8位大于1位
			if l2 <= 16 && l2 >= 6 { //强制规定密码小于16位大于6位
				user := model.User{
					Name:     username,
					Password: password,
					Question: question,
					Answer:   answer,
				}

				flag, err := service.IsRepeatUsername(username)
				if err != nil {
					fmt.Println("judge repeat username err: ", err)
					tool.RespInternalError(ctx)
					return
				}

				if flag {
					tool.RespErrorWithDate(ctx, "用户名已经存在")
					return
				}

				err = service.Register(user)
				if err != nil {
					fmt.Println("register err: ", err)
					tool.RespInternalError(ctx)
					return
				}

				tool.RespSuccessful(ctx)
			} else {
				tool.RespErrorWithDate(ctx, "密码请在16位到6位之内")
				return
			}
		} else {
			tool.RespErrorWithDate(ctx, "用户名请在8位到1位之内")
			return
		}
	} else {
		tool.RespErrorWithDate(ctx, "请将信息输入完整")
		return
	}
}
