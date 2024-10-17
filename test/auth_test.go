package test

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"hospital-information-system/auth"
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	service := auth.JwtImpl{}
	token, err := service.GenerateToken(1)

	// Assert tidak ada error
	assert.NoError(t, err)

	// Validasi token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	// Assert token valid
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	// Assert claims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		assert.Equal(t, float64(1), claims["userId"])
	} else {
		t.Errorf("Token claims tidak valid")
	}
}
func TestValidateToken_ValidToken(t *testing.T) {
	service := auth.JwtImpl{}
	// Generate a valid token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userId": 123})
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	assert.NoError(t, err)

	// Validate the token
	validatedToken, err := service.ValidateToken(signedToken)

	// Assert tidak ada error
	assert.NoError(t, err)
	assert.NotNil(t, validatedToken)
	assert.True(t, validatedToken.Valid)

	// Assert claims
	if claims, ok := validatedToken.Claims.(jwt.MapClaims); ok && validatedToken.Valid {
		assert.Equal(t, float64(123), claims["userId"])
	} else {
		t.Errorf("Token claims tidak valid")
	}
}
func TestValidateToken_InvalidToken(t *testing.T) {
	service := auth.JwtImpl{}

	// Validate an invalid token
	_, err := service.ValidateToken("invalid.token.string")
	assert.Error(t, err)
}
