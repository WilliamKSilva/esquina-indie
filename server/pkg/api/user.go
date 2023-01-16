package api

import "errors"

type UserService interface {
    New (user CreateUserRequest) (*User, error) 
}

type UserRepository interface {
    CreateUser (CreateUserRequest) (*User, error) 
}

type userService struct {
    repo UserRepository 
}

func NewUserService(repo UserRepository) UserService {
    return &userService{
        repo: repo,
    }
}

func (u *userService) New(user CreateUserRequest) (*User, error) {
    if user.Name == "" {
       return nil, errors.New("user service - Name required")
    }

    if user.Email== "" {
       return nil, errors.New("user service - Email required")
    }

    if user.Password== "" {
       return nil, errors.New("user service - Password required")
    }

    createdUser, err := u.repo.CreateUser(user)

    if err != nil {
        return nil, err 
    }

    return createdUser, nil 
}
