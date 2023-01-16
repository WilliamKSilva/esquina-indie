package web 

import (
	"net/http"

	"github.com/WilliamKSilva/esquina-indie/pkg/app"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
    UserService app.UserService
}

func (u *UserHandler) NewUser (c echo.Context) error {
    newUserData := app.NewUserData{}
    err := c.Bind(&newUserData) 

    if err != nil {
        return c.String(http.StatusBadRequest, "Bad request")
    }

    user, err := u.UserService.NewUser(newUserData)

    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
    }

    c.JSON(http.StatusOK, user) 
    
    return nil
}
