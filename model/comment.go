package model

import "time"

type Comment struct {
	Id            int       `json:"id"`
	PostID        int       `json:"post_id"`
	CommentNum    int       `json:"comment_num"`
	Txt           string    `json:"txt"`
	PidName       string    `json:"pid_name"`
	Username      string    `json:"username"`
	PostTime      time.Time `json:"post_time"`
	UpdateTime    time.Time `json:"update_time"`
	CommentStatus int       `json:"comment_status"`
	NameStatus    int       `json:"name_status"`
	PCommentId    int       `json:"p_comment_id"`
}
