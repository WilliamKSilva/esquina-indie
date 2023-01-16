package app

import "github.com/labstack/echo/v4"

type Routes struct {
    Router *echo.Echo
}

func (r *Routes) SetupRoutes(userHandler *UserHandler) {
    r.Router.POST("/users", userHandler.NewUser)
}
