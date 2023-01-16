package api

import "errors"

type UserService interface {
    New (user CreateUserRequest) error
}

type UserRepository interface {
    CreateUser (CreateUserRequest) error 
}

type userService struct {
    db UserRepository 
}

func NewUserService(repo UserRepository) UserService {
    return &userService{
        db: repo,
    }
}

func (u *userService) New(user CreateUserRequest) error {
    if user.Name == "" {
       return errors.New("user service - Name required")
    }

    if user.Email== "" {
       return errors.New("user service - Email required")
    }

    if user.Password== "" {
       return errors.New("user service - Password required")
    }

    err := u.db.CreateUser(user)

    if err != nil {
        return err
    }

    return nil
}
