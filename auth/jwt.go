package auth

import "github.com/dgrijalva/jwt-go"

type Jwt interface {
	GenerateToken(id int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
