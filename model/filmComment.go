package model

import "time"

type FilmComment struct {
	Id         int       `json:"id"`
	MovieId    int       `json:"MovieId"`
	Context    string    `json:"context"`
	Name       string    `json:"name"`
	PostTime   time.Time `json:"post_time"`
	CommentNum int       `json:"comment_num"`
	StarNum    int       `json:"star_num"`
	Likes      int       `json:"likes"`
}

type FilmCommentDetail struct {
	FilmComment
	FilmCommentReplys []FilmCommentReply
}
