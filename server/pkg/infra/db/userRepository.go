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
        ID: 0,
        Name: user.Name,
        Email: user.Email,
        Password: user.Password,
    }

    err := db.connection.Create(&createdUser).Error 
    
    if err != nil {
        return nil, err
    }

    return &createdUser, err
}

func (db UserRepositoryDb) FindUser (id int) (*app.User, error) {
    user := app.User{} 

    err := db.connection.First(&user, id).Error

    if err != nil {
        return nil, err
    }

    return &user, nil
}
