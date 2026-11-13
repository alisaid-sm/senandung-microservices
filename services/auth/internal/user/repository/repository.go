package repository

import (
	"context"
	"database/sql"

	"github.com/alisaid-sm/senandung-microservices/services/auth/internal/user/model"
)

type UserRepository interface {
	Create(ctx context.Context, u *model.User) error
}

type pgUserRepository struct {
	db *sql.DB
}

func NewPGUserRepository(db *sql.DB) UserRepository {
	return &pgUserRepository{db: db}
}

func (r *pgUserRepository) Create(ctx context.Context, u *model.User) error {
	query := `INSERT INTO users (id, email, password_hash, name, role) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, u.ID, u.Email, u.PasswordHash, u.Name, u.Role)

	return err
}
