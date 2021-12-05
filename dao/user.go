package dao

import "message-board-demo/model"

func InsertUser(user model.User) error {
	_, err := DB.Exec("INSERT INTO user(username, password)  values(?, ?);", user.Username, user.Password)
	return err
}
func UpdateUser(user model.User) error {
	_, err := DB.Exec("update user set password=? where username =?", user.Password, user.Username)
	return err
}

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("SELECT id, password FROM user where username = ? ", username)
	if row.Err() != nil {
		return user, row.Err()
	}

	err := row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
