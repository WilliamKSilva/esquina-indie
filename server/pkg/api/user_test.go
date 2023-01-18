package api

import (
	"errors"
	"testing"
	"time"
)

type dummyUserRepository struct{}

func (u dummyUserRepository) CreateUser(user NewUserData) (*User, error) {
	return &User{
		ID:        1,
		Name:      "test",
		Email:     "teste",
		Password:  "teste",
		CreatedAt: time.Now(),
	}, nil
}

func (u dummyUserRepository) FindUser(id int) (*User, error) {
	return &User{
		ID:        1,
		Name:      "test",
		Email:     "teste",
		Password:  "teste",
		CreatedAt: time.Now(),
	}, nil
}

func (u dummyUserRepository) FindUserByEmail(email string) (*User, error) {
	return &User{
		ID:        1,
		Name:      "test",
		Email:     "teste",
		Password:  "teste",
		CreatedAt: time.Now(),
	}, nil
}

func TestNewUser(t *testing.T) {
	userService := NewUserService(dummyUserRepository{})

	want := errors.New("user service - Name required").Error()
	_, got := userService.NewUser(NewUserData{Email: "test", Password: "test"})

	if got.Error() != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}
