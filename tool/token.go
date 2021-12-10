package tool

import (
	"github.com/dgrijalva/jwt-go"
	"message-board-demo/model"
	"time"
)

var MySecret = []byte("ddzYYDS")

const TokenExpireDuration = time.Hour * 2

func GenToken(username string) (string, error) {
	c := model.MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "message_board",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}
