package handler

import (
	"posts/middleware"
	"posts/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService    *service.UserService
	postService    *service.PostService
	commentService *service.CommentService
	likeService    *service.LikeService
	followService  *service.FollowerService
}

func NewHandler(userService *service.UserService, postService *service.PostService, commentService *service.CommentService, likeService *service.LikeService, followService *service.FollowerService) *Handler {
	return &Handler{
		userService:    userService,
		postService:    postService,
		commentService: commentService,
		likeService:    likeService,
		followService:  followService,
	}
}

func Run(h *Handler) *gin.Engine {
	router := gin.Default()

	// AUTH ROUTES
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", h.Login)
		authRoutes.POST("/signup", h.SignUp)
	}

	router.Use(middleware.AuthMiddleware())

	// USER ROUTES
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", h.GetAllUsers)
		userRoutes.GET("/:id", h.GetUserById)
		userRoutes.PUT("/:id", h.UpdateUser)
		userRoutes.DELETE("/:id", h.DeleteUser)
	}

	// POST ROUTES
	postRoutes := router.Group("/posts")
	{
		postRoutes.POST("/", h.CreatePost)
		postRoutes.GET("/feed", h.GetFeedPosts)
		postRoutes.GET("/user/:id", h.GetUserPosts)
		postRoutes.GET("/:id", h.GetPostById)
		postRoutes.PUT("/:id", h.UpdatePost)
		postRoutes.DELETE("/:id", h.DeletePost)
	}

	// COMMENT ROUTES
	commentRoutes := router.Group("/comments")
	{
		commentRoutes.POST("/", h.CreateComment)
		commentRoutes.GET("/:postId", h.GetCommentsByPostId)
		commentRoutes.PUT("/", h.UpdateComment)
		commentRoutes.DELETE("/:id", h.DeleteComment)
	}

	// LIKE ROUTES
	likeRoutes := router.Group("/likes")
	{
		likeRoutes.POST("/toggle", h.ToggleLike)
		likeRoutes.GET("/users", h.GetLikedUsers)
	}

	// FOLLOW ROUTES
	router.POST("/followers", h.AddFollower)
	router.DELETE("/followers", h.RemoveFollower)
	router.GET("/followers/:userId", h.GetFollowers)
	router.GET("/following/:userId", h.GetFollowing)

	return router
}
