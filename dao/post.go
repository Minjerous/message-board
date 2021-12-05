package dao

import (
	"message-board-demo/model"
)

func InsertPost(post model.Post) error {
	_, err := DB.Exec("INSERT INTO post(username, txt, post_time, update_time) "+"values(?, ?, ?, ?);", post.Username, post.Txt, post.PostTime, post.UpdateTime)
	return err
}

func SelectPosts() ([]model.Post, error) {
	var posts []model.Post
	rows, err := DB.Query("SELECT id, username, txt, post_time, update_time FROM post")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var post model.Post

		err = rows.Scan(&post.Id, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
func DeletePost(post model.Post) error {
	_, err := DB.Exec("delete from  post where  id=? and username=?", post.Id, post.Username)
	return err
}
func SelectUsernameById(id string) (model.Post, error) {
	post := model.Post{}
	row := DB.QueryRow("SELECT username,id FROM post where id=? ", id)
	if row.Err() != nil {
		return post, row.Err()
	}
	err := row.Scan(&post.Username, &post.Id)
	if err != nil {
		return post, err
	}
	return post, nil
}
