package service

import (
	"database/sql"
	"message-board-demo/dao"
	"message-board-demo/model"
)

func AddNormalCommentAtPost(comment model.Comment) error {
	err := dao.InsertNormalComment(comment)
	if err != nil {
		return err
	}
	dao.AddCommentNumByPost(comment)
	return err
}
func AddNormalCommentAtComment(comment model.Comment) error {
	err := dao.InsertNormalComment(comment)
	if err != nil {
		return err
	}
	err = dao.AddCommentNumByComment(comment)
	if err != nil {
		return err
	}
	err = dao.AddCommentNumByPost(comment)
	if err != nil {
		return err
	}
	return err
}

//匿名评论

func AddAnonymousComment(comment model.Comment) error {
	err := dao.InsertAnonymousComment(comment)
	if err != nil {
		return err
	}
	dao.AddCommentNumByPost(comment)
	return err
}

func GetComment(PostId int) ([]model.Comment, error) {
	comments, err := dao.SelectComments(PostId)
	return comments, err
}

//

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
