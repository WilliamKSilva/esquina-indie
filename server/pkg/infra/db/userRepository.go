package infra

import (
	"github.com/WilliamKSilva/esquina-indie/pkg/api"
	"gorm.io/gorm"
)

type UserRepository interface {
    CreateUser (user api.CreateUserRequest) (*api.User, error)
}

type UserRepositoryDb struct {
    connection *gorm.DB
}

func NewUserRepository (db *gorm.DB) UserRepository {
    return &UserRepositoryDb{
        connection: db,
    }
}

func (db UserRepositoryDb) CreateUser (user api.CreateUserRequest) (*api.User, error) {
    createdUser := api.User{
        Name: user.Name,
        Email: user.Email,
        Password: user.Password,
    }

    err := db.connection.Create(&createdUser).Error 
    
    if err != nil {
        return nil, err
    }

    return nil, err
}
