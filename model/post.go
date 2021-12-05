package model

import "time"

type Post struct {
	Id         int       `json:"id"`
	CommentNum int       `json:"comment_num"`
	Txt        string    `json:"txt"`
	Username   string    `json:"username"`
	PostTime   time.Time `json:"post_time"`
	UpdateTime time.Time `json:"update_time"`
}
