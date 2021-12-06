package service

import (
	"message-board-demo/dao"
	"message-board-demo/model"
)

func AddComment(comment model.Comment) error {
	err := dao.InsertComment(comment)
	if err != nil {
		return err
	}
	dao.AddCommentNum(comment)
	return err
}
