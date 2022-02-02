package model

type User struct {
	Id               int
	Name             string
	Password         string
	Question         string
	Answer           string
	SelfIntroduction string
}

type User2 struct {
	Id               int
	Name             string
	SelfIntroduction string
}
