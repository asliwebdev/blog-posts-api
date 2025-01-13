package main

import (
	"context"
	"log"
	"posts/handler"
	"posts/postgres"
	"posts/repository"
	"posts/service"

	"github.com/redis/go-redis/v9"
)

func InitRedis(ctx context.Context) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return redisClient
}

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	redisClient := InitRedis(ctx)

	followRepo := repository.NewFollowerRepo(db)
	likeRepo := repository.NewLikeRepo(db)
	commentRepo := repository.NewCommentRepo(db)

	userService := service.NewUserService(repository.NewUserRepo(db), followRepo)
	postService := service.NewPostService(repository.NewPostRepo(db), likeRepo, commentRepo)
	commentService := service.NewCommentService(commentRepo)
	likeService := service.NewLikeService(likeRepo)
	followService := service.NewFollowerService(followRepo)

	h := handler.NewHandler(userService, postService, commentService, likeService, followService)

	r := handler.Run(h, redisClient)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
