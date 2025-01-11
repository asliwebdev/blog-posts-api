package handler

import (
	"net/http"
	"posts/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Create a new comment
// @Description Create a comment for a specific post
// @Tags comments
// @Accept application/json
// @Produce application/json
// @Param request body models.CreateComment true "Comment body"
// @Success 201 {object} models.MessageResp
// @Failure 400 {object} models.ErrResp "Invalid input"
// @Failure 500 {object} models.ErrResp "Failed to create comment"
//
//	@Security		BearerAuth
//
// @Router /comments [post]
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

	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully"})
}

// @Summary Get comments by post ID
// @Description Retrieve all comments associated with a post
// @Tags comments
// @Produce application/json
// @Param postId path string true "Post ID"
// @Success 200 {array} models.Comment
// @Failure 400 {object} models.ErrResp "Invalid post ID"
// @Failure 500 {object} models.ErrResp "Failed to fetch comments"
//
//	@Security		BearerAuth
//
// @Router /comments/{postId} [get]
func (h *Handler) GetCommentsByPostId(c *gin.Context) {
	postId, err := uuid.Parse(c.Param("postId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post Id"})
		return
	}

	comments, err := h.commentService.GetCommentsByPostId(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// @Summary Update a comment
// @Description Update the content of an existing comment
// @Tags comments
// @Accept application/json
// @Produce application/json
// @Param request body models.UpdateCommentSwag true "Updated comment data"
// @Success 200 {object} models.UpdateComment
// @Failure 400 {object} models.ErrResp "Invalid input"
// @Failure 500 {object} models.ErrResp "Failed to update comment"
//
//	@Security		BearerAuth
//
// @Router /comments [put]
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

// DeleteComment deletes a comment
// @Summary Delete a comment
// @Description Delete a comment by its ID
// @Tags comments
// @Produce application/json
// @Param id path string true "Comment ID"
// @Success 200 {object} models.MessageResp "Comment deleted successfully"
// @Failure 400 {object} models.ErrResp "Invalid comment ID"
// @Failure 403 {object} models.ErrResp "Not authorized"
// @Failure 500 {object} models.ErrResp "Failed to delete comment"
// @Security BearerAuth
// @Router /comments/{id} [delete]
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
