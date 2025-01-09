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
}

func NewHandler(userService *service.UserService, postService *service.PostService, commentService *service.CommentService, likeService *service.LikeService) *Handler {
	return &Handler{
		userService:    userService,
		postService:    postService,
		commentService: commentService,
		likeService:    likeService,
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
		commentRoutes.GET("/:post_id", h.GetCommentsByPostId)
		commentRoutes.PUT("/", h.UpdateComment)
		commentRoutes.DELETE("/:id", h.DeleteComment)
	}

	// LIKE ROUTES
	LikeRoutes := router.Group("/likes")
	{
		LikeRoutes.POST("/toggle", h.ToggleLike)
		LikeRoutes.GET("/users", h.GetLikedUsers)
	}

	return router
}
