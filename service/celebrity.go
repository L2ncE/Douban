package service

import (
	"douban/dao"
	"douban/model"
)

// GetCelebrityById 通过id得到影人
func GetCelebrityById(Id int) (model.Celebrity, error) {
	return dao.SelectCelebrityById(Id)
}
