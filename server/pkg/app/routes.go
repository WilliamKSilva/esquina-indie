package app

import "github.com/labstack/echo/v4"

type Router struct {
    echo *echo.Echo
}

func (r *Router) setupRoutes(userHandler *UserHandler) {
    r.echo.POST("/users", userHandler.NewUser)
}
