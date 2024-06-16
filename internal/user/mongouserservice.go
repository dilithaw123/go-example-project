package user

import "context"

type MongoUserService struct{}

func NewMongoUserService() *MongoUserService {
	return &MongoUserService{}
}

func (s *MongoUserService) GetUserByID(ctx context.Context, id int) (*User, error) {
	return &User{ID: id, Name: "John Doe", Age: 30}, nil
}
