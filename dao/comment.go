package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	sqlStr := "update comment set name_status=1 where  id=? and username=?"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(comment.Id, comment.Username)
	//_, err := DB.Exec("update comment set name_status=1 where  id=? and username=?", comment.Id, comment.Username)
	return err
}

func SelectComments(PostId int) ([]model.Comment, error) {
	var Comments []model.Comment
	sqlStr := "select post_id,  id , username,   txt,      comment_num ,post_time, update_time ,comment_status , name_status ,pid_name FROM comment where post_id=?"
	Stmt, err := DB.Prepare(sqlStr)
	rows, err := Stmt.Query(PostId)
	//rows, err := DB.Query("select post_id,  id , username,   txt,      comment_num ,post_time, update_time ,comment_status , name_status ,pid_name FROM comment where post_id=?", PostId)
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
	_, err = Stmt.Exec(comment.Id)
	//_, err := DB.Exec("update  comment  set  comment_num=comment_num+ 1  where id = ? ;", comment.Id)
	fmt.Print("11111")
	return err
}
func GetOneComment(id int) ([]model.Comment, error) {
	var comment []model.Comment
	sqlStr := "select post_id,  id , username,   txt,      comment_num ,post_time, update_time ,comment_status , name_status ,pid_name FROM comment where post_id=?"
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

	err = rows.Scan(&Comment.PostID, &Comment.Id, &Comment.Username, &Comment.Txt, &Comment.CommentNum, &Comment.PostTime, &Comment.UpdateTime, &Comment.CommentStatus, &Comment.NameStatus, &Comment.PidName)

	if err != nil {
		return comment, err
	}

	if Comment.CommentStatus == 1 {
		Comment.Txt = "该评论已删除"
	}
	if Comment.NameStatus == 1 {
		Comment.Username = "匿名用户"
	}
	comment = append(comment, Comment)
	return comment, nil
}

//返回ChildComment中父类的信息

func GetFatherComment(id int) (*model.ChildComment, error) {
	row := DB.QueryRow("select txt, username  from comment where id = ?", id)
	var ChildComment model.ChildComment

	err := row.Scan(&ChildComment.Ptxt, &ChildComment.Puser)
	if err != nil {
		return nil, err
	}
	childComment := new(model.ChildComment)
	childComment.Pid = id
	childComment.Ptxt = ChildComment.Ptxt
	childComment.Puser = ChildComment.Puser

	return childComment, nil
}

//递归打印所有pid_comment 为当前指定的 id 所有Comment 即子类

func CirChildNodeComment(ctx *gin.Context, Pcomment *model.ChildComment) error {
	//查询所有pid 该pid的信息
	var Comment model.ChildComment
	sqlStr := "select id, txt, username, post_time,comment_num from comment where pid_comment= ?"
	Stmt, err := DB.Prepare(sqlStr)
	rows, err := Stmt.Query(Pcomment.CommentId)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Comment.CommentId, &Comment.CommentUser, &Comment.CommentTxt, &Comment.CommentTime, &Comment.CommentNum)
		if err != nil {
			return err
		}
		ChildComment := new(model.ChildComment)
		ChildComment.Pid = Comment.CommentId
		ChildComment.Ptxt = Comment.CommentTxt
		ChildComment.Puser = Comment.CommentUser

		if ChildComment.CommentId == Pcomment.CommentId {
			continue
		}
		ctx.JSON(200, gin.H{
			"Pid         ": Pcomment.Pid,
			"Puser       ": Pcomment.Puser,
			"Ptxt        ": Pcomment.Ptxt,
			"CommentId   ": Comment.CommentId,
			"CommentUser ": Comment.CommentUser,
			"CommentTxt  ": Comment.CommentTxt,
			"CommentTime ": Comment.CommentTime,
			"CommentNum  ": Comment.CommentNum,
		})

		//递归处理
		err = CirChildNodeComment(ctx, ChildComment)
		if err != nil {
			return err
		}
	}
	return nil
}
