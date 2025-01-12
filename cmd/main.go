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
	likeRepo := repository.NewLikeRepo(db)
	commentRepo := repository.NewCommentRepo(db)

	userService := service.NewUserService(repository.NewUserRepo(db), followRepo)
	postService := service.NewPostService(repository.NewPostRepo(db), likeRepo, commentRepo)
	commentService := service.NewCommentService(commentRepo)
	likeService := service.NewLikeService(likeRepo)
	followService := service.NewFollowerService(followRepo)

	h := handler.NewHandler(userService, postService, commentService, likeService, followService)

	r := handler.Run(h)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
