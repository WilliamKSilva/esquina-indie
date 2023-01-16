package main

import (
	"log"
	"net/http"

	"github.com/WilliamKSilva/esquina-indie/pkg/app"
    "github.com/WilliamKSilva/esquina-indie/pkg/web"
	"github.com/WilliamKSilva/esquina-indie/pkg/infra/db"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func run () {
    e := echo.New()
    db := connectDB()
    server := http.Server{
        Addr: ":3333",
        Handler: e,
    }

    routes := web.Routes{
        Router: e, 
    }
    
    userRepository := infra.NewUserRepository(db) 
    userService := app.NewUserService(userRepository)
    userHandler := web.UserHandler{
        UserService: userService,
    }
    routes.SetupRoutes(&userHandler)

    err := server.ListenAndServe()

    if err != http.ErrServerClosed {
        log.Fatal(err)
    }
}

func connectDB () *gorm.DB {  
    dsn := "test"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(err)
    }

    return db
}

