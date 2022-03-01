package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func SignJWTToken(key string, claims jwt.MapClaims) (string, error) {
	if claims == nil {
		logObj.LogError("SignJWTToken", "Error", "", key, "", GetTimeToString(time.Now()))
		return "", errors.New("claims is nil")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}
func ParseJWTToken(key string, token string) (jwt.MapClaims, error) {
	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logObj.LogError("ParseJWTToken", "Error", "", key, "", GetTimeToString(time.Now()))
			return nil, errors.New("unexpected signing method")
		}
		return []byte(key), nil
	})
	if err != nil {
		logObj.LogError("ParseJWTToken", "Error", "", err.Error(), "", GetTimeToString(time.Now()))
	}

	if claims, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {
		return claims, nil
	}
	return nil, errors.New("ParseJWTToken " + key)
}
