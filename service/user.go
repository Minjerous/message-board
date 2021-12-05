package service

import (
	"database/sql"
	"message-board-demo/dao"
	"message-board-demo/model"
	"message-board-demo/tool"
)

func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	if tool.Match(user.Password, password) {
		return true, nil
	} else {
		return false, nil
	}
	return true, nil
}

func IsRepeatUsername(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, err
}

func Register(user model.User) error {
	err := dao.InsertUser(user)
	return err
}
func Password(user model.User) error {
	err := dao.UpdateUser(user)
	return err
}
