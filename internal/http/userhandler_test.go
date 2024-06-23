package http

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dilithaw123/go-example-project/internal/user"
)

/*
	You can simplify this by using one stub for the whole thing
	type StubUserService struct {
		FailCalls bool
		GetUserByIDCalled bool
		user.UserService
	}

	func (s *StubUserService) GetUserByID(ctx context.Context, id int) (*user.User, error) {
		s.GetUserByIDCalled = true
		if s.FailCalls {
			return nil, errors.New("error")
		}
		return &user.User{ID: id, Name: "test", Age: 20}, nil
	}
*/

type FailUserService struct {
	GetUserByIDCalled bool
	user.UserService
}

func (s *FailUserService) GetUserByID(ctx context.Context, id int) (*user.User, error) {
	s.GetUserByIDCalled = true
	return nil, errors.New("error")
}

func (s *FailUserService) GetUsers(ctx context.Context) ([]user.User, error) {
	return nil, errors.New("error")
}

type SuccessUserService struct {
	user.UserService
}

func (s *SuccessUserService) GetUserByID(ctx context.Context, id int) (*user.User, error) {
	return &user.User{ID: id, Name: "test", Age: 20}, nil
}

func (s *SuccessUserService) GetUsers(ctx context.Context) ([]user.User, error) {
	return []user.User{
		{ID: 1, Name: "test", Age: 20},
		{ID: 2, Name: "test2", Age: 30},
	}, nil
}

func TestGetUserByIDError(t *testing.T) {
	userService := &FailUserService{}
	env := &HandlerEnv{
		UserService: userService,
		Logger:      slog.Default(),
	}
	mux := http.NewServeMux()
	env.Route(mux)
	testServer := httptest.NewServer(mux)
	defer testServer.Close()
	resp, err := http.Get(testServer.URL + "/api/user?id=1")
	if err != nil {
		t.Fatalf("http.Get failed: %v", err)
	}
	if !userService.GetUserByIDCalled { // Won't you know this is called anyway?
		t.Error("GetUserByID() not called")
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}
}

func TestGetUserByIDSuccess(t *testing.T) {
	userService := &SuccessUserService{}
	env := &HandlerEnv{
		UserService: userService,
		Logger:      slog.Default(),
	}
	mux := http.NewServeMux()
	env.Route(mux)
	testServer := httptest.NewServer(mux)
	defer testServer.Close()
	resp, err := http.Get(testServer.URL + "/api/user?id=1")
	if err != nil {
		t.Fatalf("http.Get failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
	var user user.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		t.Fatalf("json.Decode failed: %v", err)
	}
	if user.ID != 1 || user.Name != "test" || user.Age != 20 {
		t.Errorf("expected user {ID: 1, Name: test, Age: 20}, got %v", user)
	}
}

func TestGetUsersError(t *testing.T) {
	userService := &FailUserService{}
	env := &HandlerEnv{
		UserService: userService,
		Logger:      slog.Default(),
	}
	mux := http.NewServeMux()
	env.Route(mux)
	testServer := httptest.NewServer(mux)
	defer testServer.Close()
	resp, err := http.Get(testServer.URL + "/api/users")
	if err != nil {
		t.Fatalf("http.Get failed: %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}
}

func TestGetUsersSuccess(t *testing.T) {
	userService := &SuccessUserService{}
	env := &HandlerEnv{
		UserService: userService,
		Logger:      slog.Default(),
	}
	mux := http.NewServeMux()
	env.Route(mux)
	testServer := httptest.NewServer(mux)
	defer testServer.Close()
	resp, err := http.Get(testServer.URL + "/api/users")
	if err != nil {
		t.Fatalf("http.Get failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
	var users []user.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		t.Fatalf("json.Decode failed: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}
