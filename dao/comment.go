package dao

import (
	"message-board-demo/model"
)

func InsertComment(comment model.Comment) error {
	_, err := DB.Exec("INSERT INTO comment(post_id,username, txt, post_time, update_time) "+"values(?,?, ?, ?, ?);", comment.PostID, comment.Username, comment.Txt, comment.PostTime, comment.UpdateTime)
	return err
}

func SelectComments() ([]model.Comment, error) {
	var Comments []model.Comment
	rows, err := DB.Query("select post_id,  id , username,       txt,      comment_num ,post_time, update_time FROM comment")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var Comment model.Comment

		err = rows.Scan(&Comment.PostID, &Comment.Id, &Comment.Username, &Comment.Txt, &Comment.CommentNum, &Comment.PostTime, &Comment.UpdateTime)
		if err != nil {
			return nil, err
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

func DeleteComment(comment model.Comment) error {
	_, err := DB.Exec("delete from  comment where  id=? and username=?", comment.Id, comment.Username)
	return err
}
