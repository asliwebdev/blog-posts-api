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
}

func NewHandler(userService *service.UserService, postService *service.PostService, commentService *service.CommentService) *Handler {
	return &Handler{
		userService:    userService,
		postService:    postService,
		commentService: commentService,
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
	comments := router.Group("/comments")
	{
		comments.POST("/", h.CreateComment)
		comments.GET("/:post_id", h.GetCommentsByPostId)
		comments.PUT("/", h.UpdateComment)
		comments.DELETE("/:id", h.DeleteComment)
	}

	return router
}
