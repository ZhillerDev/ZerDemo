package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type customClaim struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

var (
	secret        = []byte("hellojwt")
	expTime int64 = 30
)

func GenerateToken(name string) (string, error) {
	claim := customClaim{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Unix() + expTime),
			Issuer:    "zhiller",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(secret)

}

func ParseToken(token string) (*jwt.Token, *customClaim, error) {
	claim := &customClaim{}
	t, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	fmt.Printf("parse the token: %v\n", t)
	return t, claim, err
}

func ValidateToken(token string) bool {
	_, _, err := ParseToken(token)
	if err != nil {
		if tp, ok := err.(*jwt.ValidationError); ok {
			if tp.Errors&jwt.ValidationErrorExpired != 0 {
				return false
			}
		}
	}

	return true
}
