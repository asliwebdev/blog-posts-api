package handler

import (
	"net/http"
	"posts/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatePost godoc
//
//	@Summary		Create a new post
//	@Description	Creates a new post with the provided details.
//	@Tags			posts
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		models.CreatePost	true	"Post creation payload"
//	@Success		201		{object}	models.MessageResp	"Post created successfully"
//	@Failure		400		{object}	models.ErrResp		"Invalid input"
//	@Failure		500		{object}	models.ErrResp		"Internal server error"
//	@Security		BearerAuth
//	@Router			/posts [post]
func (h *Handler) CreatePost(c *gin.Context) {
	var post models.CreatePost

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.postService.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

// GetPostById godoc
//
//	@Summary		Get a post by ID
//	@Description	Fetches the details of a post using its unique ID.
//	@Tags			posts
//	@Produce		application/json
//	@Param			id		path		string	true	"Post ID (UUID)"
//	@Success		200		{object}	models.Post		"Post details"
//	@Failure		400		{object}	models.ErrResp	"Invalid post ID"
//	@Failure		404		{object}	models.ErrResp	"Post not found"
//	@Failure		500		{object}	models.ErrResp	"Internal server error"
//	@Security		BearerAuth
//	@Router			/posts/{id} [get]
func (h *Handler) GetPostById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := h.postService.GetPostById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if post == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetFeedPosts godoc
//
//	@Summary		Get feed posts
//	@Description	Fetches posts for the authenticated user's feed.
//	@Tags			posts
//	@Produce		application/json
//	@Success		200		{array}		models.PostWithoutCounts		"List of feed posts"
//	@Failure		500		{object}	models.ErrResp	"Internal server error"
//	@Security		BearerAuth
//	@Router			/posts/feed [get]
func (h *Handler) GetFeedPosts(c *gin.Context) {
	userId := c.GetString("userId")

	posts, err := h.postService.GetFeedPosts(uuid.MustParse(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetUserPosts godoc
//
//	@Summary		Get a user's posts
//	@Description	Fetches all posts created by a specific user.
//	@Tags			posts
//	@Produce		application/json
//	@Param			id		path		string	true	"User ID (UUID)"
//	@Success		200		{array}		models.PostWithoutCounts		"List of user's posts"
//	@Failure		400		{object}	models.ErrResp	"Invalid user ID"
//	@Failure		500		{object}	models.ErrResp	"Internal server error"
//	@Security		BearerAuth
//	@Router			/posts/user/{id} [get]
func (h *Handler) GetUserPosts(c *gin.Context) {
	idParam := c.Param("id")
	userId, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	posts, err := h.postService.GetUserPosts(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// UpdatePost godoc
//
//	@Summary		Update a post
//	@Description	Updates the details of a post using its unique ID.
//	@Tags			posts
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id		path		string			true	"Post ID (UUID)"
//	@Param			request	body		models.UpdatePost	true	"Post update payload"
//	@Success		200		{object}	models.MessageResp	"Post updated successfully"
//	@Failure		400		{object}	models.ErrResp	"Invalid input or post ID"
//	@Failure		404		{object}	models.ErrResp	"Post not found"
//	@Failure		500		{object}	models.ErrResp	"Internal server error"
//	@Security		BearerAuth
//	@Router			/posts/{id} [put]
func (h *Handler) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	post.Id = id
	err = h.postService.UpdatePost(&post)
	if err != nil {
		if err.Error() == "post not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

// DeletePost godoc
//
//	@Summary		Delete a post
//	@Description	Deletes a post using its unique ID.
//	@Tags			posts
//	@Produce		application/json
//	@Param			id		path		string	true	"Post ID (UUID)"
//	@Success		200		{object}	models.MessageResp	"Post deleted successfully"
//	@Failure		400		{object}	models.ErrResp	"Invalid post ID"
//	@Failure		404		{object}	models.ErrResp	"Post not found"
//	@Failure		500		{object}	models.ErrResp	"Internal server error"
//	@Security		BearerAuth
//	@Router			/posts/{id} [delete]
func (h *Handler) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	err = h.postService.DeletePost(id)
	if err != nil {
		if err.Error() == "post not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
