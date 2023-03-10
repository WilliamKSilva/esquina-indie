package api 

type UserService interface {
    NewUser (user NewUserData) (*User, error) 
    FindUser (id int) (*User, error)
    FindUserByEmail (email string) (*User, error)
}

type UserRepository interface {
    CreateUser (user NewUserData) (*User, error)
    FindUser (id int) (*User, error)
    FindUserByEmail(email string) (*User, error)
}
    
type NewUserData struct {
    Name string `json:"name"`
    Password string `json:"password"`
    Email string `json:"email"` 
}
