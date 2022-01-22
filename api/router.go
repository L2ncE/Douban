package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()
	engine.Use(CORS())

	engine.POST("/register", register)       //注册
	engine.POST("/login", login)             //登陆
	engine.POST("/mibao", mibao)             //密保
	engine.POST("/mibao/question", question) //查询密保问题
	engine.GET("/brief1", briefMovies1)
	engine.GET("/brief2", briefMovies2)
	engine.GET("/brief3", briefMovies3)

	movieGroup := engine.Group("/movie")
	{
		movieGroup.GET("/:movie_id", movieDetail) //电影页
	}

	userGroup := engine.Group("/user")
	userGroup.Use(CORS())
	{
		userGroup.Use(JWTAuth)                      //需要token
		userGroup.POST("/password", changePassword) //修改密码
	}

	topicGroup := engine.Group("/topic")
	topicGroup.Use(CORS())
	{
		topicGroup.GET("/movie/:movie_id", briefTopics) //查看一部电影全部话题概略
		topicGroup.GET("/:topic_id", topicDetail)       //查看一条话题详细信息和其下属评论
		{
			topicGroup.Use(JWTAuth)                        //需要token
			topicGroup.POST("/:movie_id", addTopic)        //发布新话题
			topicGroup.DELETE("/:topic_id", deleteTopic)   //删除话题
			topicGroup.GET("/likes/:topic_id", topicLikes) //给话题点赞
		}
	}

	commentGroup := engine.Group("/comment")
	commentGroup.Use(CORS())
	{
		commentGroup.POST("/anonymity/:topic_id", addCommentAnonymity) //匿名评论
		{
			commentGroup.Use(JWTAuth)                            //需要token
			commentGroup.POST("/:topic_id", addComment)          //发送评论
			commentGroup.DELETE("/:comment_id", deleteComment)   //删除评论
			commentGroup.GET("/likes/:comment_id", commentLikes) //给评论点赞
		}
	}

	shortCommentGroup := engine.Group("/shortcomment")
	shortCommentGroup.Use(CORS())
	{
		shortCommentGroup.GET("/movie/:movie_id", briefShortComment) //查看一部电影全部短评
		{
			shortCommentGroup.Use(JWTAuth)                                      //需要token
			shortCommentGroup.POST("/:movie_id", addShortComment)               //发布新短评
			shortCommentGroup.DELETE("/:shortcomment_id", deleteShortComment)   //删除短评
			shortCommentGroup.GET("/likes/:shortcomment_id", shortCommentLikes) //给短评点赞
		}
		err := engine.Run()
		if err != nil {
			return
		}
	}
}
