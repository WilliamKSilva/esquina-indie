package app

import (
	"net/http"

	"github.com/WilliamKSilva/esquina-indie/pkg/api"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
    UserService api.UserService
}

func (u *UserHandler) NewUser (c echo.Context) error {
    createUserRequest := api.CreateUserRequest{}
    err := c.Bind(&createUserRequest) 

    if err != nil {
        return c.String(http.StatusBadRequest, "Bad request")
    }

    err = u.UserService.New(createUserRequest)

    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
    }
    
    return nil
}
