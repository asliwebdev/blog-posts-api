package handler

import (
	"net/http"
	"posts/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary      Toggle like
// @Description  Toggle like on a post or a comment
// @Tags         likes
// @Accept       application/json
// @Produce      application/json
// @Param        request  body      models.ToggleLikeRequest  true  "Toggle Like Request"
// @Success      200      {object}  models.MessageResp         "Like toggled successfully"
// @Failure      400      {object}  models.ErrResp        "Invalid request payload"
// @Failure      500      {object}  models.ErrResp         "Failed to toggle like"
// @Security     BearerAuth
// @Router       /likes/toggle [post]
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

	if req.PostId != uuid.Nil && req.CommentId != uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You should provide only one Id post or comment not both"})
		return
	}

	err := h.likeService.ToggleLike(uuid.MustParse(userId), req.PostId, req.CommentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to toggle like: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like toggled successfully"})
}

// @Summary      Get liked users
// @Description  Retrieve a list of users who liked a post or a comment
// @Tags         likes
// @Accept       application/json
// @Produce      application/json
// @Param        postId     query     string  false  "Post ID"
// @Param        commentId  query     string  false  "Comment ID"
// @Success      200        {array}   models.UserResponse    "List of users who liked"
// @Failure      400        {object}  models.ErrResp  "Invalid request parameters"
// @Failure      500        {object}  models.ErrResp  "Failed to get users who liked"
// @Security     BearerAuth
// @Router       /likes/users [get]
func (h *Handler) GetLikedUsers(c *gin.Context) {
	postIdStr := c.DefaultQuery("postId", "")
	commentIdStr := c.DefaultQuery("commentId", "")

	var postId, commentId uuid.UUID
	var err error

	if postIdStr == "" && commentIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You should provide Post or Comment Id as a query param"})
		return
	}

	if postIdStr != "" && commentIdStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You should provide only one query Id post or comment not both"})
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

	c.JSON(http.StatusOK, users)
}
