package handler

import (
	"posts/middleware"
	"posts/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{
		userService: userService,
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

	return router
}
