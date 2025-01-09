package handler

import (
	"net/http"
	"posts/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) ToggleLike(c *gin.Context) {
	userId := c.GetString("userId")
	var req models.ToggleLikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if req.PostId == uuid.Nil && req.CommentId == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Either PostID or CommentID must be provided"})
		return
	}

	err := h.likeService.ToggleLike(uuid.MustParse(userId), req.PostId, req.CommentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to toggle like: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like toggled successfully"})
}

func (h *Handler) GetLikedUsers(c *gin.Context) {
	postIdStr := c.DefaultQuery("postId", "")
	commentIdStr := c.DefaultQuery("commentId", "")

	var postId, commentId uuid.UUID
	var err error

	if postIdStr == "" && commentIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You should provide Post or Comment Id as a query param"})
		return
	}

	if postIdStr != "" {
		postId, err = uuid.Parse(postIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid postId"})
			return
		}
	}

	if commentIdStr != "" {
		commentId, err = uuid.Parse(commentIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid commentId"})
			return
		}
	}

	users, err := h.likeService.GetLikedUsers(postId, commentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users who liked"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
