package dao

import "message-board-demo/model"

func InsertComment(comment model.Comment) error {
	_, err := DB.Exec("INSERT INTO comment(username, txt, post_time, update_time) "+"values(?, ?, ?, ?);", comment.Username, comment.Txt, comment.PostTime, comment.UpdateTime)
	return err
}
