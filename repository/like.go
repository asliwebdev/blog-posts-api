package repository

import (
	"database/sql"
	"fmt"
	"posts/models"

	"github.com/google/uuid"
)

type LikeRepo struct {
	db *sql.DB
}

func NewLikeRepo(db *sql.DB) *LikeRepo {
	return &LikeRepo{db: db}
}

func (r *LikeRepo) AddLike(like *models.Like) error {
	query := `
		INSERT INTO likes (id, user_id, post_id, comment_id)
		VALUES ($1, $2, $3, $4)`

	if like.PostId != uuid.Nil && like.CommentId != uuid.Nil {
		return fmt.Errorf("a like cannot be associated with both a post and a comment")
	}

	_, err := r.db.Exec(query, like.Id, like.UserId, like.PostId, like.CommentId)
	return err
}

func (r *LikeRepo) CountLikes(postId, commentId uuid.UUID) (int, error) {
	query := `
		SELECT COUNT(*) 
		FROM likes 
		WHERE post_id = $1 OR comment_id = $2`

	var count int
	err := r.db.QueryRow(query, postId, commentId).Scan(&count)
	return count, err
}

func (r *LikeRepo) DeleteLike(userID, postID, commentID uuid.UUID) error {
	query := `
		DELETE FROM likes
		WHERE user_id = $1 AND (post_id = $2 OR comment_id = $3)`

	_, err := r.db.Exec(query, userID, postID, commentID)
	return err
}

func (r *LikeRepo) CheckIfLiked(userId, postId, commentId uuid.UUID) (bool, error) {
	query := `
		SELECT COUNT(*) 
		FROM likes 
		WHERE user_id = $1 AND (post_id = $2 OR comment_id = $3)`

	var count int
	err := r.db.QueryRow(query, userId, postId, commentId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *LikeRepo) GetLikedUsers(postId, commentId uuid.UUID) ([]models.UserResponse, error) {
	var users []models.UserResponse

	query := `
		SELECT u.id, u.username, u.email
		FROM likes l
		JOIN users u ON u.id = l.user_id
		WHERE l.post_id = $1 OR l.comment_id = $2`

	rows, err := r.db.Query(query, postId, commentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(&user.Id, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
