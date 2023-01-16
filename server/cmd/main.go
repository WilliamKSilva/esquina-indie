package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func run () {
    e := echo.New()
    server := http.Server{
        Addr: ":3333",
        Handler: e,
    }

    err := server.ListenAndServe()

    if err != http.ErrServerClosed {
        log.Fatal(err)
    }
}
