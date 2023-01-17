package main

import (
	"log"
	"net/http"

	"github.com/WilliamKSilva/esquina-indie/pkg/app"
	"github.com/WilliamKSilva/esquina-indie/pkg/infra/db"
	"github.com/WilliamKSilva/esquina-indie/pkg/web"
	"github.com/labstack/echo/v4"
)

func main () {
    e := echo.New()
    db := infra.ConnectDB()

    routes := web.Routes{
        Router: e, 
    }
    
    userRepository := infra.NewUserRepository(db) 
    userService := app.NewUserService(userRepository)
    userHandler := web.UserHandler{
        UserService: userService,
    }

    routes.SetupRoutes(&userHandler)

    err := e.Start(":3333") 
    if err != http.ErrServerClosed {
        log.Fatal(err)
    }
}



