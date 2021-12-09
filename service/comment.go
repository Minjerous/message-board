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
	//增加评论的数量
	err = dao.AddCommentNumByComment(comment)
	if err != nil {
		return err
	}
	//增加文章的评论数量
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
func GetOneComment(id int) ([]model.Comment, error) {
	Comments, err := dao.GetOneComment(id)
	return Comments, err
}
func SelectCommentsByPidComment(PComment int) ([]model.Comment, error) {
	Comments, err := dao.SelectCommentsByPidComment(PComment)
	return Comments, err
}

//func GetPid(id int) (*model.ChildComment, error) {
//	ChildComment, err := dao.GetFatherComment(id)
//	return ChildComment, err
//}
//
//func CirCommentPrint(ctx *gin.Context, PComent *model.ChildComment) error {
//	err := dao.CirChildNodeComment(ctx, PComent)
//	return err
//}
