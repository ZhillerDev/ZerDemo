package test

import (
	"fmt"
	"ginmod/src/middleware"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := middleware.GenerateToken("jack")
	if err != nil {
		return
	}
	fmt.Printf("%v\n", token)
}

func TestParseToken(t *testing.T) {
	originToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiamFjayIsImV4cCI6NX0.eDuTXjPYsQ9S8X78vNsIqEFCIe3aIdaj2mXvp3eqRjI"
	_, claim, _ := middleware.ParseToken(originToken)
	fmt.Printf("%v\n", claim)
}

func TestValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiamFjayIsImV4cCI6MTY3NDM2OTc5NSwiaXNzIjoiemhpbGxlciJ9.3uQ9DpIOP16Z-hiOqJr5hqBA51-ZzQwhvStprAZg2oo"
	fmt.Println(middleware.ValidateToken(token))
}
