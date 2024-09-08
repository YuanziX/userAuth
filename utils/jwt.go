package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yuanzix/userAuth/internal/database"
	"github.com/yuanzix/userAuth/models"
)

func CreateToken(auth database.Auth) (string, error) {
	claims := jwt.MapClaims{
		"email":     auth.UserEmail,
		"auth_uuid": auth.AuthUuid,
		"iss":       "userAuth",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey, err := ReadJWTSecret()
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(r *http.Request, checkExists func(models.AuthDetails) (bool, error)) (string, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return "", err
	}

	auth, err := ExtractTokenAuth(r)
	if err != nil {
		return "", err
	}

	exists, err := checkExists(*auth)
	if err != nil {
		return "", err
	}

	if !exists || !token.Valid {
		return "", errors.New("invalid token")
	}

	return auth.UserEmail, nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractTokenString(r)
	if tokenString == "" {
		return nil, errors.New("token not provided")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, err
		}
		return ReadJWTSecret()
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractTokenAuth(r *http.Request) (*models.AuthDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		authUUIDStr, ok := claims["auth_uuid"].(string)
		if !ok {
			return nil, errors.New("invalid auth_uuid claim")
		}

		authUuid, err := uuid.Parse(authUUIDStr)
		if err != nil {
			return nil, errors.New("failed to parse auth_uuid")
		}

		userEmail, ok := claims["email"].(string)
		if !ok {
			return nil, errors.New("invalid email claim")
		}

		return &models.AuthDetails{
			AuthUUID:  authUuid,
			UserEmail: userEmail,
		}, nil
	}

	return nil, errors.New("invalid token claims")
}

func ExtractTokenString(r *http.Request) string {
	token := r.URL.Query().Get("token")
	if token != "" {
		return token
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString != "" {
		strArr := strings.Split(tokenString, " ")
		if len(strArr) == 2 {
			return strArr[1]
		}
	}

	return ""
}
