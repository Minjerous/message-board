package dao

import "message-board-demo/model"

func InsertUser(user model.User) error {
	sqlStr := "INSERT INTO user(username, password)  values(?, ?);"
	Stmt, err := DB.Prepare(sqlStr)
	Stmt.Exec(user.Username, user.Password)
	return err
}

func UpdateUser(user model.User) error {
	sqlStr := "update user set password=? where username =?"
	Stmt, err := DB.Prepare(sqlStr)
	Stmt.Exec(user.Password, user.Username)
	return err
}

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	sqlStr := "SELECT id, password FROM user where username = ? "
	Stmt, err := DB.Prepare(sqlStr)
	row := Stmt.QueryRow(username)
	if row.Err() != nil {
		return user, row.Err()
	}
	err = row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
