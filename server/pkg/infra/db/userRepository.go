package infra

import (
	"github.com/WilliamKSilva/esquina-indie/pkg/api"
	"gorm.io/gorm"
)

type UserRepository interface {
    Create(user api.CreateUserRequest) (api.User, error)
}

type Db struct {
    connection *gorm.DB
}

func (db Db) Create(user api.CreateUserRequest) (*api.User, error) {
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
