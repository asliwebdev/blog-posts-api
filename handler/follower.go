package handler

import (
	"net/http"
	"posts/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary      Add a follower
// @Description  Adds a follower to a user
// @Tags         followers
// @Accept       application/json
// @Produce      application/json
// @Param        request body models.FollowRequest true "Follow Request"
// @Success      200  {object} models.MessageResp "message: Follower added"
// @Failure      400  {object} models.ErrResp "error: Invalid request"
// @Failure      500  {object} models.ErrResp "error: Failed to add follower"
// @Security	 BearerAuth
// @Router       /followers [post]
func (h *Handler) AddFollower(c *gin.Context) {
	var req models.FollowRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.followService.AddFollower(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add follower"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Follower added"})
}

// @Summary      Remove a follower
// @Description  Removes a follower from a user
// @Tags         followers
// @Accept       application/json
// @Produce      application/json
// @Param        follower_id query string true "Follower ID"
// @Param        following_id query string true "Following ID"
// @Success      200  {object} models.MessageResp "message: Follower removed"
// @Failure      400  {object} models.ErrResp "error: Missing or invalid parameters"
// @Failure      500  {object} models.ErrResp "error: Failed to remove follower"
// @Security	 BearerAuth
// @Router       /followers [delete]
func (h *Handler) RemoveFollower(c *gin.Context) {
	followerId := c.Query("follower_id")
	followingId := c.Query("following_id")

	if followerId == "" || followingId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing follower_id or following_id in query"})
		return
	}

	followerUUID, err1 := uuid.Parse(followerId)
	followingUUID, err2 := uuid.Parse(followingId)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := h.followService.RemoveFollower(followerUUID, followingUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove follower"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Follower removed"})
}

// @Summary      Get followers
// @Description  Retrieves the list of followers for a user
// @Tags         followers
// @Produce      application/json
// @Param        userId path string true "User ID"
// @Success      200  {array} models.UserResponse
// @Failure      400  {object} models.ErrResp "error: Invalid userId"
// @Failure      500  {object} models.ErrResp "error: Failed to fetch followers"
// @Security	 BearerAuth
// @Router       /followers/{userId} [get]
func (h *Handler) GetFollowers(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId"})
		return
	}

	followers, err := h.followService.GetFollowers(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch followers"})
		return
	}

	c.JSON(http.StatusOK, followers)
}

// @Summary      Get following
// @Description  Retrieves the list of users a user is following
// @Tags         followers
// @Produce      application/json
// @Param        userId path string true "User ID"
// @Success      200  {array} models.UserResponse
// @Failure      400  {object} models.ErrResp "error: Invalid userId"
// @Failure      500  {object} models.ErrResp "error: Failed to fetch following"
// @Security	 BearerAuth
// @Router       /following/{userId} [get]
func (h *Handler) GetFollowing(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId"})
		return
	}

	following, err := h.followService.GetFollowing(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch following"})
		return
	}

	c.JSON(http.StatusOK, following)
}
