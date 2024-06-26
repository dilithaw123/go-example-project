package user

import (
	"context"
	"database/sql"

	"github.com/georgysavva/scany/v2/sqlscan"
)

type DBUserService struct {
	db *sql.DB
}

func NewDBUserService(db *sql.DB) *DBUserService {
	return &DBUserService{db: db}
}

func (s *DBUserService) GetUserByID(ctx context.Context, id int) (*User, error) {
	var user User
	err := sqlscan.Get(ctx, s.db, &user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *DBUserService) GetUsers(ctx context.Context) ([]User, error) {
	var users []User
	err := sqlscan.Select(ctx, s.db, &users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}
