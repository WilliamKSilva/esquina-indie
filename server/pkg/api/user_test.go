package api

import (
	"errors"
	"testing"
	"time"
)

type dummyUserRepository struct {}


func (u dummyUserRepository) CreateUser(user NewUserData) (*User, error) {
	return &User{
		ID:        1,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
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
	return nil , nil
}

func TestNewUserShouldThrowIfNameMissing(t *testing.T) {
	userService := NewUserService(dummyUserRepository{})

	want := errors.New("user service - Name required").Error()
	_, got := userService.NewUser(NewUserData{Email: "test", Password: "test"})

	if got.Error() != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}

func TestNewUserShouldThrowIfEmailMissing(t *testing.T) {
	userService := NewUserService(dummyUserRepository{})

	want := errors.New("user service - Email required").Error()
	_, got := userService.NewUser(NewUserData{Name: "test", Password: "test"})

	if got.Error() != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}

func TestNewUserShouldThrowIfPasswordMissing(t *testing.T) {
	userService := NewUserService(dummyUserRepository{})

	want := errors.New("user service - Password required").Error()
	_, got := userService.NewUser(NewUserData{Email: "test", Name: "test"})

	if got.Error() != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}

func TestNewUserShouldCallRepositoryAndReturnAUser(t *testing.T) {
    userRepositorySpy := &dummyUserRepository{}
	userService := NewUserService(userRepositorySpy)
	newUserData := NewUserData{Email: "test@test.com", Name: "test", Password: "test12345"}

	user, _:= userService.NewUser(newUserData)

    if user.Name != newUserData.Name {
        t.Errorf("Want '%s', got '%s'", newUserData.Name, user.Name)
    }
}
