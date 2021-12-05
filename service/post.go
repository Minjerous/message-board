package service

import (
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
