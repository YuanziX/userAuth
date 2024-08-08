package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/yuanzix/userAuth/models"
	"github.com/yuanzix/userAuth/utils"
)

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) (statusCode int, err error) {
	users, err := s.store.GetAllUsers()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return utils.WriteJSON(w, http.StatusOK, models.DatabaseUsersToUserResponses(users))
}

func (s *APIServer) handleGetUserByEmail(w http.ResponseWriter, r *http.Request) (statusCode int, err error) {
	email := r.PathValue("email")

	user, err := s.store.GetUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}
	return utils.WriteJSON(w, http.StatusOK, models.DatabaseUserToUserResponse(user))
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) (statusCode int, err error) {
	type parameters struct {
		Email       string `json:"email"`
		Username    string `json:"username"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Password    string `json:"password"`
		DateOfBirth string `json:"date_of_birth"`
		Somethingso string `json:"something"`
	}

	params := parameters{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		return http.StatusBadRequest, err
	}

	dob, err := utils.StringDateToTimeObject(params.DateOfBirth)
	if err != nil {
		return http.StatusBadRequest, err
	}

	hashedPassword, err := utils.HashPassword(params.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user := models.User{
		Email:          params.Email,
		Username:       params.Username,
		HashedPassword: hashedPassword,
		FirstName:      params.FirstName,
		LastName:       params.LastName,
		DateOfBirth:    dob,
	}

	databaseUser, err := s.store.CreateUser(&user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return utils.WriteJSON(w, http.StatusCreated, models.DatabaseUserToUserResponse(databaseUser))
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) (statusCode int, err error) {
	email := r.PathValue("email")

	if err = s.store.DeleteUser(email); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]string{"deleted": email})
}
