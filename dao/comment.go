package dao

import (
	"fmt"
	"message-board-demo/model"
)

//我觉得很憨  但我不想改了

func InsertNormalComment(comment model.Comment) error {
	sqlStr := "INSERT INTO comment(post_id,username, txt, post_time, update_time,pid_name,pid_comment) values(?,?, ?, ?, ?,?,?);"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(comment.PostID, comment.Username, comment.Txt, comment.PostTime, comment.UpdateTime, comment.PidName, comment.PCommentId)
	//_, err := DB.Exec("INSERT INTO comment(post_id,username, txt, post_time, update_time,pid_name) "+"values(?,?, ?, ?, ?,?);", comment.PostID, comment.Username, comment.Txt, comment.PostTime, comment.UpdateTime, comment.PidName)
	return err
}

func InsertAnonymousComment(comment model.Comment) error {
	InsertNormalComment(comment)
	sqlStr := "update comment set name_status='ture' where  id=? and username=?"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(comment.Id, comment.Username)
	//_, err := DB.Exec("update comment set name_status=1 where  id=? and username=?", comment.Id, comment.Username)
	return err
}

func SelectComments(PostId int) ([]model.Comment, error) {
	var Comments []model.Comment
	sqlStr := "select post_id,  id , username,   txt,      comment_num ,post_time, update_time ,comment_status , name_status ,pid_name,pid_comment FROM comment where post_id=?"
	Stmt, err := DB.Prepare(sqlStr)
	rows, err := Stmt.Query(PostId)
	//rows, err := DB.Query("select post_id,  id , username,   txt,      comment_num ,post_time, update_time ,comment_status , name_status ,pid_name FROM comment where post_id=?", PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var Comment model.Comment

		err = rows.Scan(&Comment.PostID, &Comment.Id, &Comment.Username, &Comment.Txt, &Comment.CommentNum, &Comment.PostTime, &Comment.UpdateTime, &Comment.CommentStatus, &Comment.NameStatus, &Comment.PidName, &Comment.PCommentId)
		if err != nil {
			return nil, err
		}
		if Comment.CommentStatus == 1 {
			Comment.Txt = "该评论已删除"
		}
		if Comment.NameStatus == "true" {
			Comment.Username = "匿名用户"
		}

		Comments = append(Comments, Comment)
	}

	return Comments, nil
}

func SelectUsernameByIdByComment(id string) (model.Comment, error) {
	comment := model.Comment{}

	sqlStr := "SELECT username,id FROM comment where id=? "
	Stmt, err := DB.Prepare(sqlStr)
	row := Stmt.QueryRow(id)
	//row := DB.QueryRow("SELECT username,id FROM comment where id=? ", id)
	if row.Err() != nil {
		return comment, row.Err()
	}
	err = row.Scan(&comment.Username, &comment.Id)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

//func DeleteComment(comment model.Comment) error {
//	_, err := DB.Exec("delete from  comment where  id=? and username=?", comment.Id, comment.Username)
//	return err
//}

//将comment_status改为一并跟新时间

func DeleteComment(comment model.Comment) error {
	sqlStr := "update comment set update_time=? ,comment_status=1 where  id=? and username=?"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(comment.UpdateTime, comment.Id, comment.Username)
	//_, err := DB.Exec("update comment set update_time=? ,comment_status=1 where  id=? and username=?", comment.UpdateTime, comment.Id, comment.Username)
	return err
}

func SelectCommentUserByIdByComment(Id int) (string, error) {
	var Username string
	sqlStr := "SELECT username FROM comment where id=? "
	Stmt, err := DB.Prepare(sqlStr)
	row := Stmt.QueryRow(Id)
	//row := DB.QueryRow("SELECT username FROM comment where id=? ", Id)
	if row.Err() != nil {
		return Username, row.Err()
	}
	err = row.Scan(&Username)
	if err != nil {
		return Username, err
	}
	return Username, nil
}

//通过 comment 的id 来更新 comment_num

func AddCommentNumByComment(comment model.Comment) error {
	sqlStr := "update  comment  set  comment_num=comment_num+ 1  where id = ? ;"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(comment.PCommentId)
	//_, err := DB.Exec("update  comment  set  comment_num=comment_num+ 1  where id = ? ;", comment.Id)
	fmt.Print("11111")
	return err
}
func GetOneComment(id int) ([]model.Comment, error) {
	var comment []model.Comment
	sqlStr := "select post_id,  id , username,   txt,      comment_num ,post_time, update_time ,comment_status , name_status ,pid_name ,pid_comment FROM comment where id=?"
	Stmt, err := DB.Prepare(sqlStr)
	rows := Stmt.QueryRow(id)
	//rows, err := DB.Query("select post_id,  id , username,   txt,      comment_num ,post_time, update_time ,comment_status , name_status ,pid_name FROM comment where post_id=?", PostId)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return comment, rows.Err()
	}

	var Comment model.Comment

	err = rows.Scan(&Comment.PostID, &Comment.Id, &Comment.Username, &Comment.Txt, &Comment.CommentNum, &Comment.PostTime, &Comment.UpdateTime, &Comment.CommentStatus, &Comment.NameStatus, &Comment.PidName, &Comment.PCommentId)

	if err != nil {
		return comment, err
	}

	if Comment.CommentStatus == 1 {
		Comment.Txt = "该评论已删除"
	}
	if Comment.NameStatus == "true" {
		Comment.Username = "匿名用户"
	}
	comment = append(comment, Comment)
	return comment, nil
}
func SelectCommentsByPidComment(PComment int) ([]model.Comment, error) {
	var Comments []model.Comment
	sqlStr := "select post_id,  id , username,   txt,      comment_num ,post_time, update_time ,comment_status , name_status ,pid_name,pid_comment FROM comment where pid_comment=?"
	Stmt, err := DB.Prepare(sqlStr)
	rows, err := Stmt.Query(PComment)
	//rows, err := DB.Query("select post_id,  id , username,   txt,      comment_num ,post_time, update_time ,comment_status , name_status ,pid_name FROM comment where post_id=?", PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var Comment model.Comment

		err = rows.Scan(&Comment.PostID, &Comment.Id, &Comment.Username, &Comment.Txt, &Comment.CommentNum, &Comment.PostTime, &Comment.UpdateTime, &Comment.CommentStatus, &Comment.NameStatus, &Comment.PidName, &Comment.PCommentId)
		if err != nil {
			return nil, err
		}

		if Comment.CommentStatus == 1 {
			Comment.Txt = "该评论已删除"
		}
		if Comment.NameStatus == "true" {
			Comment.Username = "匿名用户"
		}
		Comments = append(Comments, Comment)
	}
	return Comments, nil
}
