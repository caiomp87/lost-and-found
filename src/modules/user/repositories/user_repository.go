package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/caiomp87/lost-and-found/src/modules/user/entities"
)

type IUserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindAll(ctx context.Context) []*entities.User
	FindByID(ctx context.Context, id int64) *entities.User
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	query := fmt.Sprintf("insert into %s (name, email, password_hash) values (?, ?, ?)", user.TableName())

	_, err := r.DB.ExecContext(ctx, query, user.Name, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindAll(ctx context.Context) []*entities.User {
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) *entities.User {
	return nil
}
