package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"posts/models"
	"time"

	"github.com/google/uuid"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (p *PostRepo) CreatePost(post *models.Post) error {
	id := uuid.NewString()

	query := `INSERT INTO posts (id, user_id, title, content) VALUES ($1, $2, $3, $4)`
	_, err := p.db.Exec(query, id, post.UserId, post.Title, post.Content)

	return err
}

func (p *PostRepo) GetPostById(id uuid.UUID) (*models.Post, error) {
	post := &models.Post{}
	query := `SELECT id, user_id, title, content, created_at, updated_at FROM posts WHERE id = $1`
	err := p.db.QueryRow(query, id).Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error querying post by id: %w", err)
	}

	return post, nil
}

func (p *PostRepo) GetUserPosts(userId uuid.UUID) ([]models.Post, error) {
	query := `SELECT id, user_id, title, content, created_at, updated_at FROM posts WHERE user_id = $1`
	rows, err := p.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *PostRepo) GetFeedPosts(userId uuid.UUID) ([]models.Post, error) {
	query := `SELECT id, user_id, title, content, created_at, updated_at 
              FROM posts 
              WHERE user_id != $1 
              ORDER BY created_at DESC`
	rows, err := p.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepo) UpdatePost(post *models.Post) error {
	query := `UPDATE posts SET title = $1, content = $2, updated_at = $3 WHERE id = $4`
	result, err := r.db.Exec(query, post.Title, post.Content, time.Now(), post.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("post not found")
	}

	return nil
}

func (r *PostRepo) DeletePost(id uuid.UUID) error {
	query := `DELETE FROM posts WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("post not found")
	}

	return nil
}
