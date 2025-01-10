package handler

import (
	"net/http"
	"posts/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) AddFollower(c *gin.Context) {
	var req models.FollowRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.followService.AddFollower(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add follower"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Follower added"})
}

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

	c.JSON(http.StatusOK, gin.H{"followers": followers})
}

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
