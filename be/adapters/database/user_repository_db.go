package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/peesaphanthavong/core/ports"
	"github.com/peesaphanthavong/models"
	"github.com/peesaphanthavong/models/response/user"
)

type UserRepositoryDB struct {
	pool *pgxpool.Pool
}

func NewUserRepositoryDB(pool *pgxpool.Pool) ports.UserRepository {
	return &UserRepositoryDB{
		pool: pool,
	}
}

func (r *UserRepositoryDB) GetAllUser(ctx context.Context) ([]user.UserResponse, error) {
	var users = []user.UserResponse{}
	query := `select username,firstname,lastname,email from users`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user = user.UserResponse{}
		err = rows.Scan(&user.UserName, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	log.Printf("GetAllUser: %d users found", len(users))
	return users, nil
}

func (r *UserRepositoryDB) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (username, firstname, lastname, email, password) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.pool.Exec(ctx, query, user.UserName, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
