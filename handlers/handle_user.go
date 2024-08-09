package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

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
		if strings.Contains(err.Error(), "duplicate key") {
			return http.StatusConflict, errors.New("the email is already registered")
		}
		return http.StatusInternalServerError, err
	}

	return utils.WriteJSON(w, http.StatusCreated, models.DatabaseUserToUserResponse(databaseUser))
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) (statusCode int, err error) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := parameters{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		return http.StatusBadRequest, err
	}

	hashedPassword, err := s.store.GetHashedPassword(params.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusUnauthorized, errors.New("incorrect email or password")
		}
		return http.StatusInternalServerError, err
	}

	err = utils.CompareHashAndPassword(hashedPassword, params.Password)
	if err != nil {
		return http.StatusUnauthorized, errors.New("incorrect email or password")
	}

	auth, err := s.store.CreateAuth(params.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	tokenString, err := utils.SignIn(*auth)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return utils.WriteJSON(w, http.StatusAccepted, map[string]string{"login": "successful", "token_string": tokenString})
}

func (s *APIServer) handleLogout(w http.ResponseWriter, r *http.Request) (statusCode int, err error) {
	statusCode, err = s.deleteAuth(r)
	if err != nil {
		return statusCode, err
	}
	return utils.WriteJSON(w, http.StatusAccepted, map[string]string{"logged_out": "successfully"})
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) (statusCode int, err error) {
	email := r.PathValue("email")

	statusCode, err = s.deleteAuth(r)
	if err != nil {
		return statusCode, err
	}

	if err = s.store.DeleteUser(email); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]string{"deleted": email})
}

func (s *APIServer) deleteAuth(r *http.Request) (statusCode int, err error) {
	auth, err := utils.ExtractTokenAuth(r)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	err = s.store.DeleteAuth(*auth)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
