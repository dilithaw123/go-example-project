package user

import (
	"context"
	"log/slog"
)

type MongoUserService struct {
	Logger *slog.Logger
}

func NewMongoUserService(log *slog.Logger) *MongoUserService {
	return &MongoUserService{
		Logger: log.With("service", "MongoUserService"),
	}
}

var users = []User{
	{ID: 1, Name: "John Doe", Age: 30},
	{ID: 2, Name: "Jane Doe", Age: 25},
	{ID: 3, Name: "John Smith", Age: 40},
	{ID: 4, Name: "Jane Smith", Age: 35},
}

func (s *MongoUserService) GetUserByID(ctx context.Context, id int) (*User, error) {
	s.Logger.Debug("Getting user by id", "id", id)
	return &User{ID: id, Name: "John Doe", Age: 30}, nil
}

func (s *MongoUserService) GetUsers(ctx context.Context) ([]User, error) {
	s.Logger.Debug("Getting all users")
	return users, nil
}
