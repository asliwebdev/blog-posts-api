package handler

import (
	"net/http"
	"posts/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateComment(c *gin.Context) {
	userId := c.GetString("userId")
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.UserId = uuid.MustParse(userId)

	if err := h.commentService.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func (h *Handler) GetCommentsByPostId(c *gin.Context) {
	postId, err := uuid.Parse(c.Param("post_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post Id"})
		return
	}

	comments, err := h.commentService.GetCommentsByPostId(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func (h *Handler) UpdateComment(c *gin.Context) {
	userId := c.GetString("userId")
	var comment models.UpdateComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.UserId = uuid.MustParse(userId)

	if err := h.commentService.UpdateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *Handler) DeleteComment(c *gin.Context) {
	commentId := c.Param("id")
	userId := c.GetString("userId")

	id, err := uuid.Parse(commentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment Id"})
		return
	}

	err = h.commentService.DeleteComment(id, uuid.MustParse(userId))
	if err != nil {
		if err.Error() == "comment not found or not authorized" {
			c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment deleted successfully"})
}
