package app

import (
	"errors"
	"time"
)

type userService struct {
    repo UserRepository 
}

type User struct {
    ID int `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
    CreatedAt time.Time `json:"created_at"`
}


func NewUserService(repo UserRepository) UserService {
    return &userService{
        repo: repo,
    }
}

func (u *userService) NewUser(user NewUserData) (*User, error) {
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

func (u *userService) FindUser (id int) (*User, error) {
    user, err := u.repo.FindUser(id) 

    if err != nil {
        return nil, err
    }

    return user, nil
}
