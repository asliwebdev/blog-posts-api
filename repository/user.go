package repository

import (
	"database/sql"
	"fmt"
	"posts/models"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) CreateUser(user *models.User) error {
	user.Id = uuid.New()

	_, err := u.DB.Exec(`INSERT INTO users (id, username, email, password)
	 VALUES ($1, $2, $3, $4)`,
		user.Id, user.Username, user.Email, user.Password)

	return err
}

func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := u.DB.QueryRow(`SELECT id, username, email, password, created_at, updated_at FROM users WHERE email = $1`, email).
		Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error querying user: %w", err)
	}
	return user, nil
}

func (u *UserRepo) GetUserByEmailOrUsername(email, username string) (*models.User, error) {
	user := &models.User{}
	err := u.DB.QueryRow(`
		SELECT id, username, email
		FROM users
		WHERE email = $1 OR username = $2
	`, email, username).Scan(&user.Id, &user.Username, &user.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error querying user by email or username: %w", err)
	}
	return user, nil
}

func (u *UserRepo) GetUserById(userId uuid.UUID) (*models.User, error) {
	user := &models.User{}
	err := u.DB.QueryRow(`SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = $1`, userId).
		Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error querying user by ID: %w", err)
	}
	return user, nil
}

func (u *UserRepo) GetAllUsers() ([]models.User, error) {
	rows, err := u.DB.Query(`SELECT id, username, email, password, created_at, updated_at FROM users`)
	if err != nil {
		return nil, fmt.Errorf("error querying all users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepo) UserExists(userId uuid.UUID) (bool, error) {
	var exists bool
	err := u.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`, userId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking user existence: %w", err)
	}
	return exists, nil
}

func (u *UserRepo) UpdateUser(user *models.UpdateUser) error {
	query := `UPDATE users SET username = $1, email = $2, updated_at = $3`

	updateValues := []interface{}{user.Username, user.Email, time.Now()}

	if user.Password != "" {
		query += ", password = $4"
		updateValues = append(updateValues, user.Password)
	}

	query += " WHERE id = $" + strconv.Itoa(len(updateValues)+1)
	updateValues = append(updateValues, user.Id)

	_, err := u.DB.Exec(query, updateValues...)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (u *UserRepo) DeleteUser(userId uuid.UUID) error {
	_, err := u.DB.Exec(`DELETE FROM users WHERE id = $1`, userId)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}
