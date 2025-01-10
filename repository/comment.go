package repository

import (
	"database/sql"
	"errors"
	"posts/models"
	"time"

	"github.com/google/uuid"
)

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

func (r *CommentRepo) CreateComment(comment *models.Comment) error {
	id := uuid.NewString()
	query := `
		INSERT INTO comments (id, post_id, user_id, parent_comment_id, content) 
		VALUES ($1, $2, $3, $4, $5)`

	var parentCommentId interface{}
	if comment.ParentCommentId == uuid.Nil {
		parentCommentId = nil
	} else {
		parentCommentId = comment.ParentCommentId
	}

	_, err := r.db.Exec(query, id, comment.PostId, comment.UserId, parentCommentId, comment.Content)
	return err
}

func (r *CommentRepo) GetCommentsByPostId(postId uuid.UUID) ([]models.Comment, error) {
	query := `
		SELECT 
			c.id, c.post_id, c.user_id, c.parent_comment_id, c.content, c.created_at, c.updated_at, 
			u.id, u.username, u.email,
			COALESCE((SELECT COUNT(*) FROM likes WHERE comment_id = c.id), 0) AS likes_count
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = $1`
	rows, err := r.db.Query(query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		comment := models.Comment{}
		user := models.UserResponse{}
		err := rows.Scan(
			&comment.Id, &comment.PostId, &comment.UserId, &comment.ParentCommentId, &comment.Content,
			&comment.CreatedAt, &comment.UpdatedAt,
			&user.Id, &user.Username, &user.Email,
			&comment.LikesCount,
		)
		if err != nil {
			return nil, err
		}
		comment.User = user
		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *CommentRepo) UpdateComment(comment *models.UpdateComment) error {
	query := `
        UPDATE comments 
        SET content = $1, updated_at = $2 
        WHERE id = $3 AND user_id = $4`
	result, err := r.db.Exec(query, comment.Content, time.Now(), comment.Id, comment.UserId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("comment not found or not authorized")
	}

	return nil
}

func (r *CommentRepo) DeleteComment(commentId, userId uuid.UUID) error {
	query := `DELETE FROM comments WHERE id = $1 AND user_id = $2`
	result, err := r.db.Exec(query, commentId, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("comment not found or not authorized")
	}

	return nil
}

func (r *CommentRepo) CountComments(postId uuid.UUID) (int, error) {
	query := `
		SELECT COUNT(*) 
		FROM comments 
		WHERE post_id = $1`

	var count int
	err := r.db.QueryRow(query, postId).Scan(&count)
	return count, err
}
