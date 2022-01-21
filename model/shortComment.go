package model

import "time"

type ShortComment struct {
	Id          int
	MovieId     int
	Name        string
	Status      string
	StarNum     int
	CommentTime time.Time
	Likes       int
	Context     string
}
