package user

import (
	"errors"
	"testing"
	"time"

	"github.com/WilliamKSilva/esquina-indie/pkg/api"
)

type dummyUserRepositoryFindUser struct {}

func (u dummyUserRepositoryFindUser) CreateUser(user api.NewUserData) (*api.User, error) {
	return &api.User{
		ID:        1,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	}, nil
}

func (u dummyUserRepositoryFindUser) FindUser(id int) (*api.User, error) {
	return &api.User{
		ID:        1,
		Name:      "test",
		Email:     "test@test.com",
		Password:  "test12345",
		CreatedAt: time.Now(),
	}, nil
}

func (u dummyUserRepositoryFindUser) FindUserByEmail(email string) (*api.User, error) {
	return nil , nil
}

func TestFindUserShouldReturnAnErrorIfInvalidIdIsProvided (t *testing.T) {
    dummyUserRepository := &dummyUserRepositoryFindUser{}
    userService := api.NewUserService(dummyUserRepository) 

    want := errors.New("user service - Invalid id").Error()
    _, got := userService.FindUser(0)

    if want != got.Error() {
        t.Errorf("Want '%s', got '%s'", want, got)
    }

}

func TestFindUserShoulReturnAnUserOnSuccess (t *testing.T) {
    dummyUserRepository := &dummyUserRepositoryFindUser{}
    userService := api.NewUserService(dummyUserRepository) 

    want := "test"
    got, _ := userService.FindUser(1)

    if want != got.Name {
        t.Errorf("Want '%s', got '%s'", want, got.Name)
    }
}
