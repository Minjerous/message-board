package service

import (
	"database/sql"
	"message-board-demo/dao"
	"message-board-demo/model"
)

func GetPosts() ([]model.Post, error) {
	posts, err := dao.SelectPosts()
	return posts, err
}

func AddPost(post model.Post) error {
	err := dao.InsertPost(post)
	return err
}
func DeletePost(post model.Post) error {
	err := dao.DeletePost(post)
	return err
}
func IsUsernameMachIdByPost(username, id string) (bool, error) {
	post, err := dao.SelectUsernameByIdByPost(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if username == post.Username {
		return true, nil
	} else {
		return false, nil
	}
}
