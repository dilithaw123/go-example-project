package user

import "context"

type UserService interface {
	GetUserByID(ctx context.Context, id int) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
}
