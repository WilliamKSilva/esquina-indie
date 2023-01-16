package api

type User struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Password string `json:"password"`
    Email string `json:"email"` 
}
    
type CreateUserRequest struct {
    Name string `json:"name"`
    Password string `json:"password"`
    Email string `json:"email"` 
}
