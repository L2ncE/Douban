package service

import (
	"douban/dao"
	"douban/model"
)

// GetMovieById 通过id得到电影
func GetMovieById(movieId int) (model.Movie, error) {
	return dao.SelectMovieById(movieId)
}

// GetMovies 得到电影
func GetMovies() ([]model.Movie, error) {
	return dao.SelectMovie()
}
