package model

type User struct {
	Id        int    `json:"userId"`
	Username  string `json:"nikename"`
	Password  string
	Avatar    string
	Anonymous string `json:"anonymous"`
}
