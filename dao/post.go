package dao

import (
	"fmt"
	"message-board-demo/model"
)

//预处理示范 防止sql注入 由于时间关系就其他dao层的就不更改了
func InsertPost(post model.Post) error {
	sqlStr := "INSERT INTO post(username, txt, post_time, update_time) " + "values(?, ?, ?, ?);"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post.Username, post.Txt, post.PostTime, post.UpdateTime)
	return err
}

func SelectPosts() ([]model.Post, error) {
	var posts []model.Post
	rows, err := DB.Query("SELECT id, username, txt, post_time, update_time ,comment_num FROM post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post model.Post

		err = rows.Scan(&post.Id, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime, &post.CommentNum)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func DeletePost(post model.Post) error {
	_, err := DB.Exec("delete from  post where  id=? and username=?", post.Id, post.Username)
	_, err = DB.Exec("update  post set  update_time=? where id = ? ;", post.UpdateTime, post.Id)
	return err
}

func SelectUsernameByIdByPost(id string) (model.Post, error) {
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

func AddCommentNumByPost(comment model.Comment) error {
	_, err := DB.Exec("update  post set  comment_num=comment_num+ 1  where id = ? ;", comment.PostID)
	fmt.Print("11111")
	return err
}

func MiuCommentNum(comment model.Comment) error {
	_, err := DB.Exec("update  post set comment_num=comment_num- 1 where id = ? ;", comment.PostID)
	fmt.Print("888888")
	return err
}

func SelectPostIdByIdByComment(id int) (model.Comment, error) {
	comment := model.Comment{}
	row := DB.QueryRow("SELECT post_id FROM comment where id=? ", id)
	if row.Err() != nil {
		return comment, row.Err()
	}
	err := row.Scan(&comment.PostID)
	if err != nil {
		return comment, err
	}
	return comment, nil
}
func SelectPostUserByPostIdByPost(postId int) (string, error) {
	var Username string
	row := DB.QueryRow("SELECT username FROM post where id=? ", postId)
	if row.Err() != nil {
		return Username, row.Err()
	}
	err := row.Scan(&Username)
	if err != nil {
		return Username, err
	}
	return Username, nil
}
