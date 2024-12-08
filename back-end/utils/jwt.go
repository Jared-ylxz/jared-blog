package utils

import (
	"jaredBlog/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.MapClaims
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Expires  int64  `json:"exp"`
}

func GenerateToken(userID uint, username string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		Username: username,
		Expires:  time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.AppConfig.App.Secret))
	return "Bearer " + signedToken, err
}

func ParseToken(inputToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(inputToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.App.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
