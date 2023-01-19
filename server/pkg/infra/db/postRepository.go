package infra

import (
	"github.com/WilliamKSilva/esquina-indie/pkg/api"
	"gorm.io/gorm"
)

type PostRepositoryDb struct {
	connection *gorm.DB
}

func NewPostRepository (db *gorm.DB) api.PostRepository {
    return &PostRepositoryDb{db}
}

func (db PostRepositoryDb) CreatePost(post api.NewPostData, userId int) (*api.Post, error) {
	newPost := &api.Post{
        UserId: userId,
		Title:       post.Title,
		Description: post.Description,
	}

	err := db.connection.Create(newPost).Error

	if err != nil {
		return nil, err
	}

    return newPost, nil
}
