package api

import "time"

type Post struct {
    ID int `json:"id" gorm:"primaryKey"` 
    UserId int `json:"id" gorm:"primaryKey"` 
    Title string `json:"title"` 
    Description string `json:"description"` 
    CreatedAt time.Time `json:"created_at"`
}


