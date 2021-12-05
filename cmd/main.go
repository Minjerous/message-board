package main

import (
	"message-board-demo/api"
	"message-board-demo/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
