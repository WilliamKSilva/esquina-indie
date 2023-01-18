package user 

import (
	"errors"
	"testing"
	"time"

	"github.com/WilliamKSilva/esquina-indie/pkg/api"
)

type dummyUserRepositoryCreateUser struct {}

func (u dummyUserRepositoryCreateUser) CreateUser(user api.NewUserData) (*api.User, error) {
	return &api.User{
		ID:        1,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	}, nil
}

func (u dummyUserRepositoryCreateUser) FindUser(id int) (*api.User, error) {
	return &api.User{
		ID:        1,
		Name:      "test",
		Email:     "teste",
		Password:  "teste",
		CreatedAt: time.Now(),
	}, nil
}

func (u dummyUserRepositoryCreateUser) FindUserByEmail(email string) (*api.User, error) {
	return nil , nil
}

func TestNewUserShouldThrowIfNameMissing(t *testing.T) {
	userService := api.NewUserService(dummyUserRepositoryCreateUser{})

	want := errors.New("user service - Name required").Error()
	_, got := userService.NewUser(api.NewUserData{Email: "test", Password: "test"})

	if got.Error() != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}

func TestNewUserShouldThrowIfEmailMissing(t *testing.T) {
	userService := api.NewUserService(dummyUserRepositoryCreateUser{})

	want := errors.New("user service - Email required").Error()
	_, got := userService.NewUser(api.NewUserData{Name: "test", Password: "test"})

	if got.Error() != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}

func TestNewUserShouldThrowIfPasswordMissing(t *testing.T) {
	userService := api.NewUserService(dummyUserRepositoryCreateUser{})

	want := errors.New("user service - Password required").Error()
	_, got := userService.NewUser(api.NewUserData{Email: "test", Name: "test"})

	if got.Error() != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}

func TestNewUserShouldCallRepositoryAndReturnAUser(t *testing.T) {
    userRepositorySpy := &dummyUserRepositoryCreateUser{}
	userService := api.NewUserService(userRepositorySpy)
	newUserData := api.NewUserData{Email: "test@test.com", Name: "test", Password: "test12345"}

	user, _:= userService.NewUser(newUserData)

    if user.Name != newUserData.Name {
        t.Errorf("Want '%s', got '%s'", newUserData.Name, user.Name)
    }
}
