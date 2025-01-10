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

	followRepo := repository.NewFollowerRepo(db)

	userService := service.NewUserService(repository.NewUserRepo(db), followRepo)
	postService := service.NewPostService(repository.NewPostRepo(db))
	commentService := service.NewCommentService(repository.NewCommentRepo(db))
	likeService := service.NewLikeService(repository.NewLikeRepo(db))
	followService := service.NewFollowerService(followRepo)

	h := handler.NewHandler(userService, postService, commentService, likeService, followService)

	r := handler.Run(h)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
