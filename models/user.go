package models

import (
	"time"

	"github.com/yuanzix/userAuth/internal/database"
)

type User struct {
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DateOfBirth    time.Time `json:"date_of_birth"`
}

type UserResponse struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func DatabaseUserToUserResponse(u *database.User) UserResponse {
	return UserResponse{
		Email:     u.Email,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func DatabaseUsersToUserResponses(dbUsers *[]database.User) *[]UserResponse {
	users := []UserResponse{}

	for _, dbUser := range *dbUsers {
		users = append(users, DatabaseUserToUserResponse(&dbUser))
	}

	return &users
}
