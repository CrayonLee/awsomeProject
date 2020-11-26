package common

import (
	"gin_vue/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_create")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	//这里设置过期时间为7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	cliams := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "crayon",
			Subject:   "user token",
		},
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,cliams)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "",err
	}
	return tokenString,err

}

func ParseToken(tokenString string)(*jwt.Token,*Claims,error)  {
	cliams:=&Claims{}
	token, err := jwt.ParseWithClaims(tokenString, cliams, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token,cliams,err

}
