package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", register) //注册
	engine.POST("/login", login)       //登陆

	userGroup := engine.Group("/user")
	{
		userGroup.Use(JWTAuth)                      //需要token
		userGroup.POST("/password", changePassword) //修改密码
	}

	err := engine.Run()
	if err != nil {
		return
	}
}