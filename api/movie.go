package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// movieDetail
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

func WTW(ctx *gin.Context) {
	//iUsername, _ := ctx.Get("username")
	username := "小袁"
	movieIdString := ctx.Param("movie_id") //输入电影id
	movieId, err := strconv.Atoi(movieIdString)
	URL := service.GetURLById(movieId)
	err = service.ChangeWTWId(username, movieId)
	if err != nil {
		fmt.Println("change  err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	err2 := service.ChangeWTWURL(username, URL)
	if err2 != nil {
		fmt.Println("change  err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}
