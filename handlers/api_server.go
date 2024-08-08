package handlers

import (
	"log"
	"net/http"

	"github.com/yuanzix/userAuth/utils"
)

type APIServer struct {
	listenAddress string
	store         utils.Storage
}

type apiFunc func(http.ResponseWriter, *http.Request) (statusCode int, err error)

func NewAPIServer(listenAddress string, store utils.Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	router.HandleFunc("GET /users", makeHTTPHandlerFunc(s.handleGetUsers))

	router.HandleFunc("POST /user", makeHTTPHandlerFunc(s.handleCreateUser))
	router.HandleFunc("GET /user/{email}", makeProtectedHandlerFunc(s.handleGetUserByEmail))
	router.HandleFunc("DELETE /user/{email}", makeProtectedHandlerFunc(s.handleDeleteUser))

	router.HandleFunc("POST /login", makeHTTPHandlerFunc(s.handleLogin))

	log.Printf("JSON API server running on port: %v\n", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, err := f(w, r)
		if err != nil {
			utils.WriteErrorJSON(w, code, err.Error())
		}
	}
}

func makeProtectedHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.WriteErrorJSON(w, http.StatusUnauthorized, "jwt token string not provided")
			return
		}

		if err := utils.ValidateToken(tokenString); err != nil {
			utils.WriteErrorJSON(w, http.StatusUnauthorized, "Invalid token: "+err.Error())
			return
		}

		code, err := f(w, r)
		if err != nil {
			utils.WriteErrorJSON(w, code, err.Error())
		}
	}
}
