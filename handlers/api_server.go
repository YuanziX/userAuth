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

	router.HandleFunc("GET /users", s.makeHTTPHandlerFunc(s.handleGetUsers))

	router.HandleFunc("POST /user", s.makeHTTPHandlerFunc(s.handleCreateUser))
	router.HandleFunc(" /user/{email}/verify", s.makeHTTPHandlerFunc(s.handleVerifyUser))

	router.HandleFunc("GET /user/{email}", s.makeProtectedHandlerFunc(s.handleGetUserByEmail))
	router.HandleFunc("DELETE /user/{email}", s.makeProtectedHandlerFunc(s.handleDeleteUser))

	router.HandleFunc("POST /login", s.makeHTTPHandlerFunc(s.handleLogin))
	router.HandleFunc("GET /logout", s.makeProtectedHandlerFunc(s.handleLogout))

	log.Printf("JSON API server running on port: %v\n", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, err := f(w, r)
		if err != nil {
			utils.WriteErrorJSON(w, code, err.Error())
		}
	}
}

func (s *APIServer) makeProtectedHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := utils.ValidateToken(r, s.store.CheckAuthExists); err != nil {
			utils.WriteErrorJSON(w, http.StatusUnauthorized, "Invalid token: "+err.Error())
			return
		}

		code, err := f(w, r)
		if err != nil {
			utils.WriteErrorJSON(w, code, err.Error())
		}
	}
}
