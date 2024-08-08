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

	router.HandleFunc("GET /user/{email}", makeHTTPHandlerFunc(s.handleGetUserByEmail))
	router.HandleFunc("POST /user", makeHTTPHandlerFunc(s.handleCreateUser))
	router.HandleFunc("DELETE /user/{email}", makeHTTPHandlerFunc(s.handleDeleteUser))

	log.Printf("JSON API server running on port: %v\n", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := f(w, r); err != nil {
			utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		}
	}
}
