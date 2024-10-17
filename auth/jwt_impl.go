package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type JwtImpl struct {
}

func NewJwt() Jwt {
	return &JwtImpl{}
}

func (service JwtImpl) GenerateToken(id int) (string, error) {
	payload := jwt.MapClaims{}
	payload["userId"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	SecretKey := []byte(os.Getenv("SECRET_KEY"))

	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (service JwtImpl) ValidateToken(token string) (*jwt.Token, error) {
	parse, err := jwt.Parse(token, service.parse)
	if err != nil {
		return nil, err
	}
	return parse, nil
}
func (service JwtImpl) parse(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return []byte(os.Getenv("SECRET_KEY")), nil
}
