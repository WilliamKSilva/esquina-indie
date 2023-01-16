package api

type User struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Password string `json:"password"`
    Email string `json:"email"` 
    Role Role `json:"role"` 
    Social Social `json:"social"`
}

type Role struct {
    Name string `json:"name"`
    Description string `json:"description"`
    Stack []string `json:"stack"`
}

type Social struct {
    Name string `json:"name"`
    Url string `json:"url"`
}

type CreateUserRequest struct {
    Name string `json:"name"`
    Password string `json:"password"`
    Email string `json:"email"` 
    Role Role `json:"role"` 
    Social Social `json:"social"`
}
