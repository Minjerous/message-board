package service

import (
	"database/sql"
	"message-board-demo/dao"
	"message-board-demo/model"
)

//匿名评论

func AddNormalComment(comment model.Comment) error {
	err := dao.InsertNormalComment(comment)
	if err != nil {
		return err
	}
	dao.AddCommentNum(comment)
	return err
}

//匿名评论

func AddAnonymousComment(comment model.Comment) error {
	err := dao.InsertAnonymousComment(comment)
	if err != nil {
		return err
	}
	dao.AddCommentNum(comment)
	return err
}

func GetComment() ([]model.Comment, error) {
	comments, err := dao.SelectComments()
	return comments, err
}

func IsUsernameMachIdByComment(username, id string) (bool, error) {
	comment, err := dao.SelectUsernameByIdByComment(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if username == comment.Username {
		return true, nil
	} else {
		return false, nil
	}
}

func DeleteComment(comment model.Comment) error {
	err := dao.DeleteComment(comment)
	return err
}
