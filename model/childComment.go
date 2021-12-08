package model

import "time"

type ChildComment struct {
	CommentId   int
	CommentUser string
	CommentTxt  string
	CommentTime time.Time
	CommentNum  int
	Pid         int
	Puser       string
	Ptxt        string
}
