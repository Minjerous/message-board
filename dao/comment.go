package dao

import (
	"fmt"
	"message-board-demo/model"
)

//我觉得很憨  但我不想改了

func InsertNormalComment(comment model.Comment) error {
	_, err := DB.Exec("INSERT INTO comment(post_id,username, txt, post_time, update_time,pid_name) "+"values(?,?, ?, ?, ?,?);", comment.PostID, comment.Username, comment.Txt, comment.PostTime, comment.UpdateTime, comment.PidName)
	return err
}

func InsertAnonymousComment(comment model.Comment) error {
	InsertNormalComment(comment)
	_, err := DB.Exec("update comment set name_status=1 where  id=? and username=?", comment.Id, comment.Username)
	return err
}

func SelectComments(PostId int) ([]model.Comment, error) {
	var Comments []model.Comment
	rows, err := DB.Query("select post_id,  id , username,   txt,      comment_num ,post_time, update_time ,comment_status , name_status ,pid_name FROM comment where post_id=?", PostId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var Comment model.Comment

		err = rows.Scan(&Comment.PostID, &Comment.Id, &Comment.Username, &Comment.Txt, &Comment.CommentNum, &Comment.PostTime, &Comment.UpdateTime, &Comment.CommentStatus, &Comment.NameStatus, &Comment.PidName)
		if err != nil {
			return nil, err
		}
		if Comment.CommentStatus == 1 {
			Comment.Txt = "该评论已删除"
		}
		if Comment.NameStatus == 1 {
			Comment.Username = "匿名用户"
		}
		Comments = append(Comments, Comment)
	}

	return Comments, nil
}

func SelectUsernameByIdByComment(id string) (model.Comment, error) {
	comment := model.Comment{}
	row := DB.QueryRow("SELECT username,id FROM comment where id=? ", id)
	if row.Err() != nil {
		return comment, row.Err()
	}
	err := row.Scan(&comment.Username, &comment.Id)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

//func DeleteComment(comment model.Comment) error {
//	_, err := DB.Exec("delete from  comment where  id=? and username=?", comment.Id, comment.Username)
//	return err
//}
func DeleteComment(comment model.Comment) error {
	_, err := DB.Exec("update comment set update_time=? ,comment_status=1 where  id=? and username=?", comment.UpdateTime, comment.Id, comment.Username)
	return err
}

func SelectCommentUserByIdByComment(Id int) (string, error) {
	var Username string
	row := DB.QueryRow("SELECT username FROM comment where id=? ", Id)
	if row.Err() != nil {
		return Username, row.Err()
	}
	err := row.Scan(&Username)
	if err != nil {
		return Username, err
	}
	return Username, nil
}
func AddCommentNumByComment(comment model.Comment) error {
	_, err := DB.Exec("update  comment  set  comment_num=comment_num+ 1  where id = ? ;", comment.Id)
	fmt.Print("11111")
	return err
}
