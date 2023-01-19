package web

import (
	"net/http"
	"strconv"

	"github.com/WilliamKSilva/esquina-indie/pkg/api"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
    UserService api.UserService
}

type PostHandler struct {
    PostService api.PostService
}

func (u *UserHandler) NewUser (c echo.Context) error {
    newUserData := api.NewUserData{}
    err := c.Bind(&newUserData) 

    if err != nil {
        return c.String(http.StatusBadRequest, "Bad request")
    }

    user, err := u.UserService.NewUser(newUserData)

    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

    c.JSON(http.StatusOK, user) 
    
    return nil
}

func (u *UserHandler) FindUser (c echo.Context) error {
    userId := c.Param("id") 

    userIdInt, err := strconv.Atoi(userId)

    if err != nil {
        return c.String(http.StatusBadRequest, "Invalid user id") 
    }

    user, err := u.UserService.FindUser(userIdInt)

    if err != nil {
        return c.String(http.StatusBadRequest, "User not found")
    }

    c.JSON(http.StatusOK, user)

    return nil
}

type header struct {
    Authorization string `header:"Authorization"`
}
     
func (u *PostHandler) NewPost (c echo.Context) error {
    auth := header{}
    c.Bind(&auth)

    newPostData := api.NewPostData{}

    c.Bind(&newPostData)

    userId, err := strconv.Atoi(auth.Authorization)

    if err != nil {
        return c.String(http.StatusBadRequest, "post service - Invalid authorization header")
    }

    post, err := u.PostService.NewPost(newPostData, userId)

    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

    c.JSON(http.StatusOK, post)

    return nil
}
    
