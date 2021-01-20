package services

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaim struct {
	UserName string `json:"userName`
	jwt.StandardClaims
}

func GenerateToken(claims TokenClaim) (string, error) {
	mySigningKey := []byte("WitTharit")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)

	return tokenStr, err
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("WitTharit"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
