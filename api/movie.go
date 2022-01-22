package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// movieDetail 话题详细信息和其下属评论
func movieDetail(ctx *gin.Context) {
	movieIdString := ctx.Param("movie_id") //输入电影id
	movieId, err := strconv.Atoi(movieIdString)
	if err != nil {
		fmt.Println("movie id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "movie_id格式有误")
		return
	}

	//根据movieId拿到movie
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		fmt.Println("get movie by id err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movie)
}

func briefMovies1(ctx *gin.Context) {
	movies, err := service.GetMovies1()
	if err != nil {
		fmt.Println("get movies err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movies)
}

func briefMovies2(ctx *gin.Context) {
	movies, err := service.GetMovies2()
	if err != nil {
		fmt.Println("get movies err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movies)
}

func briefMovies3(ctx *gin.Context) {
	movies, err := service.GetMovies3()
	if err != nil {
		fmt.Println("get movies err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movies)
}
