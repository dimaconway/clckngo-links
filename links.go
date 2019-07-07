package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "clckngo/links/ApiSpecification"
)

func main() {
    e := echo.New()

    var api ApiSpecification.LinksApiImplementation
    ApiSpecification.RegisterHandlers(e, &api)

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Logger.Fatal(e.Start(":8080"))
}
