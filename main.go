package main

import (
	"log"

	"github.com/yuanzix/userAuth/handlers"
	"github.com/yuanzix/userAuth/utils"
)

func main() {
	store, err1 := utils.NewPostgresStore()
	_, err2 := utils.ReadJWTSecret()
	if err1 != nil && err2 != nil {
		log.Fatal(err1, err2)
	}

	server := handlers.NewAPIServer(":3000", store)
	server.Run()
}
