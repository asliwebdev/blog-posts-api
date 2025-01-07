package main

import (
	"log"
	"posts/handler"
	"posts/postgres"
	"posts/repository"
	"posts/service"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userService := service.NewUserService(repository.NewUserRepo(db))

	h := handler.NewHandler(userService)

	r := handler.Run(h)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
