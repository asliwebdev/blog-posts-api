package repository

import (
	"database/sql"
	"posts/models"

	"github.com/google/uuid"
)

type FollowerRepo struct {
	db *sql.DB
}

func NewFollowerRepo(db *sql.DB) *FollowerRepo {
	return &FollowerRepo{db: db}
}

func (r *FollowerRepo) AddFollower(follower *models.FollowRequest) error {
	id := uuid.NewString()

	query := `
		INSERT INTO followers (id, follower_id, following_id)
		VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, id, follower.FollowerId, follower.FollowingId)
	return err
}

func (r *FollowerRepo) RemoveFollower(followerId, followingId uuid.UUID) error {
	query := `
		DELETE FROM followers 
		WHERE follower_id = $1 AND following_id = $2`
	_, err := r.db.Exec(query, followerId, followingId)
	return err
}

func (r *FollowerRepo) CountFollowers(userId uuid.UUID) (int, error) {
	query := `
		SELECT COUNT(*) 
		FROM followers 
		WHERE following_id = $1`
	var count int
	err := r.db.QueryRow(query, userId).Scan(&count)
	return count, err
}

func (r *FollowerRepo) CountFollowing(userId uuid.UUID) (int, error) {
	query := `
		SELECT COUNT(*) 
		FROM followers 
		WHERE follower_id = $1`
	var count int
	err := r.db.QueryRow(query, userId).Scan(&count)
	return count, err
}

func (r *FollowerRepo) GetFollowers(userId uuid.UUID) ([]models.UserResponse, error) {
	query := `
		SELECT u.id, u.username, u.email
		FROM users u
		JOIN followers f ON u.id = f.follower_id
		WHERE f.following_id = $1`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(&user.Id, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		followers = append(followers, user)
	}

	return followers, nil
}

func (r *FollowerRepo) GetFollowing(userId uuid.UUID) ([]models.UserResponse, error) {
	query := `
		SELECT u.id, u.username, u.email
		FROM users u
		JOIN followers f ON u.id = f.following_id
		WHERE f.follower_id = $1`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var following []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(&user.Id, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		following = append(following, user)
	}

	return following, nil
}
