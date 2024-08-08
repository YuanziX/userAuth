package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/yuanzix/userAuth/internal/database"
)

func CreateToken(user database.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"uAt": user.UpdatedAt,
		"iss": "userAuth",
	})

	secretKey, _ := ReadJWTSecret()

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
