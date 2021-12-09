package model

import "time"

type Comment struct {
	Id            int       `json:"cid"`
	PostID        int       `json:"post_id"`
	CommentNum    int       `json:"comment_num"`
	Txt           string    `json:"txt"`
	PidName       string    `json:"pid_name"`
	Username      string    `json:"nikename"`
	PostTime      time.Time `json:"post_time"`
	UpdateTime    time.Time `json:"update_time"`
	CommentStatus int
	NameStatus    string `json:"anonymous"`
	PCommentId    int    `json:"parent_id"`
}
