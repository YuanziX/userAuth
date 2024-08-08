package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"iss":   "userAuth",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey, _ := ReadJWTSecret()

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secretKey, _ := ReadJWTSecret()
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
