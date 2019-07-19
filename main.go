package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "clckngo/links/controllers"
    "clckngo/links/server"
)

func main() {
    e := echo.New()

    var api controllers.ApiImplementation
    server.RegisterHandlers(e, &api)

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Logger.Fatal(e.Start(":8080"))
}
