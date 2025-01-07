package handler

import (
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

	router.POST("/users", h.CreateUser)

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
