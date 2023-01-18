package infra

import (
	"github.com/WilliamKSilva/esquina-indie/pkg/api"
	"gorm.io/gorm"
)

type UserRepositoryDb struct {
    connection *gorm.DB
}

func NewUserRepository (db *gorm.DB) api.UserRepository {
    return &UserRepositoryDb{
        connection: db,
    }
}

func (db UserRepositoryDb) CreateUser (user api.NewUserData) (*api.User, error) {
    createdUser := api.User{
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

func (db UserRepositoryDb) FindUser (id int) (*api.User, error) {
    user := api.User{} 

    err := db.connection.First(&user, id).Error

    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (db UserRepositoryDb) FindUserByEmail (email string) (*api.User, error) {
    user := api.User{} 

    err := db.connection.First(&user, "email = ?", email).Error

    if err != nil {
        return nil, err
    }

    return &user, nil
}
