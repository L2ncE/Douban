package service

import (
	"douban/dao"
	"douban/model"
)

// GetMovieById 通过id得到电影
func GetMovieById(movieId int) (model.Movie, error) {
	return dao.SelectMovieById(movieId)
}

// GetMovies1 得到电影
func GetMovies1() ([]model.MovieBrief, error) {
	return dao.SelectMovie1()
}

func GetMovies2() ([]model.MovieBrief, error) {
	return dao.SelectMovie2()
}
func GetMovies3() ([]model.MovieBrief, error) {
	return dao.SelectMovie3()
}

// GetURLById 通过id得到URL
func GetURLById(Id int) string {
	return dao.SelectURLById(Id)
}

// GetMovie 得到电影
func GetMovie() ([]model.Rank1, error) {
	return dao.SelectMovie()
}
