package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"hospital-information-system/auth"
	"os"
	"testing"
	"time"
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
func TestDate(t *testing.T) {
	// Buat objek time.Time
	currentTime := time.Now()

	// Konversi ke string dengan format yang diinginkan
	dateString := currentTime.Format("2006-01-02") // Format tanggal
	timeString := currentTime.Format("15:04:05")   // Format waktu

	// Output hasil konversi
	fmt.Println("Tanggal:", dateString)
	fmt.Println("Waktu:", timeString)
}
