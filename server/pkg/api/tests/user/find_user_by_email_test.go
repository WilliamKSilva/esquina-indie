package user

import (
	"errors"
	"testing"
	"time"

	"github.com/WilliamKSilva/esquina-indie/pkg/api"
)

type dummyUserRepositoryFindUserByEmail struct {}

func (u dummyUserRepositoryFindUserByEmail) CreateUser(user api.NewUserData) (*api.User, error) {
	return &api.User{
		ID:        1,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	}, nil
}

func (u dummyUserRepositoryFindUserByEmail) FindUser(id int) (*api.User, error) {
	return &api.User{
		ID:        1,
		Name:      "test",
		Email:     "test@test.com",
		Password:  "test12345",
		CreatedAt: time.Now(),
	}, nil
}

func (u dummyUserRepositoryFindUserByEmail) FindUserByEmail(email string) (*api.User, error) {
	return &api.User{
		ID:        1,
		Name:      "test",
		Email:     "test@test.com",
		Password:  "test12345",
		CreatedAt: time.Now(),
	}, nil
}

func TestFindUserByEmailThrowIfEmailNotProvided (t *testing.T) {
    dummyUserRepository := dummyUserRepositoryFindUserByEmail{}
    userService := api.NewUserService(dummyUserRepository)

    want := errors.New("user service - Email required").Error() 
    _, got := userService.FindUserByEmail("")

    if want != got.Error() {
        t.Errorf("Want '%s', got '%s'", want, got)
    }
}
