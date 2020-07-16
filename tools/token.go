package tools

import (
	"gingo/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("jwt_key")

// Claims jwt
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// ReleaseToken 发放token
func ReleaseToken(user model.User) (string, error) {
	expiresTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gingo",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
