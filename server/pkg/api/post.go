package api

import (
	"errors"
	"time"
)

type Post struct {
    ID int `json:"id" gorm:"primaryKey"` 
    UserId int `json:"user_id"` 
    Title string `json:"title"` 
    Description string `json:"description"` 
    CreatedAt time.Time `json:"created_at"`
}

type postService struct {
    repo PostRepository
}

func NewPostService (repo PostRepository) *postService {
    return &postService{repo}
}

func (u postService) NewPost (post NewPostData, userId int) (*Post, error) {
    if post.Title == "" {
        return nil, errors.New("post service - Title is required")
    }

    if post.Description== "" {
        return nil, errors.New("post service - Description is required")
    }

    createdPost, err := u.repo.CreatePost(post, userId)  

    if err != nil {
        return nil, err
    }
    
    return createdPost, nil
}


