package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()
	engine.Use(cors.Default())

	engine.POST("/register", register)       //注册
	engine.POST("/login", login)             //登陆
	engine.POST("/mibao", mibao)             //密保
	engine.POST("/mibao/question", question) //查询密保问题
	engine.GET("/:movie_id", movieDetail)    //电影页
	engine.GET("/movie", briefMovies)        //所有电影

	userGroup := engine.Group("/user")
	{
		userGroup.Use(JWTAuth)                      //需要token
		userGroup.POST("/password", changePassword) //修改密码
	}

	topicGroup := engine.Group("/topic")
	{
		topicGroup.POST("/movie_id", briefTopics) //查看一部电影全部话题概略
		topicGroup.GET("/:topic_id", topicDetail) //查看一条话题详细信息和其下属评论
		{
			topicGroup.Use(JWTAuth)                         //需要token
			topicGroup.POST("/", addTopic)                  //发布新话题
			topicGroup.DELETE("/:topic_id", deleteTopic)    //删除话题
			topicGroup.POST("/:topic_id/likes", topicLikes) //给话题点赞
		}
	}

	commentGroup := engine.Group("/comment")
	{
		commentGroup.POST("/anonymity", addCommentAnonymity) //匿名评论
		{
			commentGroup.Use(JWTAuth)                             //需要token
			commentGroup.POST("/", addComment)                    //发送评论
			commentGroup.DELETE("/:comment_id", deleteComment)    //删除评论
			commentGroup.POST("/:comment_id/likes", commentLikes) //给评论点赞
		}
	}

	err := engine.Run()
	if err != nil {
		return
	}
}
