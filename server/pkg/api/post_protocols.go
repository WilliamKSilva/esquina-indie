package api

type PostService interface {
    NewPost (post NewPostData, userId int) (*Post, error)
}

type PostRepository interface {
    CreatePost (post NewPostData, userId int) (*Post, error)
}

type NewPostData struct {
    Title string `json:"title"`
    Description string `json:"description"`
}



