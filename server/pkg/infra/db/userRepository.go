package infra

import (
	"github.com/WilliamKSilva/esquina-indie/pkg/app"
	"gorm.io/gorm"
)

type UserRepositoryDb struct {
    connection *gorm.DB
}

func NewUserRepository (db *gorm.DB) app.UserRepository {
    return &UserRepositoryDb{
        connection: db,
    }
}

func (db UserRepositoryDb) CreateUser (user app.NewUserData) (*app.User, error) {
    createdUser := app.User{
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
