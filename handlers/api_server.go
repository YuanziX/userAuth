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
type apiAuthFunc func(http.ResponseWriter, *http.Request, string) (statusCode int, err error)

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
	router.HandleFunc("GET /user", s.makeProtectedHandlerFunc(s.handleGetUserByEmail))
	router.HandleFunc("DELETE /user", s.makeProtectedHandlerFunc(s.handleDeleteUser))

	router.HandleFunc("GET /user/verify", s.makeProtectedHandlerFunc(s.handleVerifyUser))
	router.HandleFunc("GET /user/isVerified", s.makeHTTPHandlerFunc(s.handleIsVerified))
	router.HandleFunc("GET /user/resendVerificationMail", s.makeHTTPHandlerFunc(s.handleResendVerificationMail))

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

func (s *APIServer) makeProtectedHandlerFunc(af apiAuthFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, err := utils.ValidateToken(r, s.store.CheckAuthExists)
		if err != nil {
			utils.WriteErrorJSON(w, http.StatusUnauthorized, "Invalid token: "+err.Error())
			return
		}

		code, err := af(w, r, email)
		if err != nil {
			utils.WriteErrorJSON(w, code, err.Error())
		}
	}
}
