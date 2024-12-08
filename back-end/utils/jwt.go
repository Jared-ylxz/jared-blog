package utils

import (
	"errors"
	"exchangeapp/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	signedToken, err := token.SignedString([]byte(config.AppConfig.App.Secret))
	return "Bearer " + signedToken, err
}

func ParseToken(inputToken string) (uint, string, error) {
	if len(inputToken) > 7 && inputToken[:7] == "Bearer " {
		inputToken = inputToken[7:]
	}

	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.AppConfig.App.Secret), nil
	})

	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userID"].(uint)
		if !ok {
			return 0, "", errors.New("userID claim is not a uint")
		}
		username, ok := claims["username"].(string)
		if !ok {
			return 0, "", errors.New("username claim is not a string")
		}
		return userID, username, nil
	}

	return 0, "", err
}
