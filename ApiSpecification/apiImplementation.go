package ApiSpecification

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

// Implements `ServerInterface`
type LinksApiImplementation struct{}

func (w *LinksApiImplementation) GetLinks(ctx echo.Context) error {
    return ctx.String(http.StatusOK, "GetLinks")
}

func (w *LinksApiImplementation) StoreLink(ctx echo.Context) error {
    return ctx.String(http.StatusOK, "StoreLink")
}

func (w *LinksApiImplementation) DeleteLink(ctx echo.Context, hash string) error {
    return ctx.String(http.StatusOK, "DeleteLink")
}

func (w *LinksApiImplementation) GetLink(ctx echo.Context, hash string) error {
    return ctx.String(http.StatusOK, "GetLink")
}

func (w *LinksApiImplementation) IncrementHitsCounter(ctx echo.Context, hash string) error {
    return ctx.String(http.StatusOK, "IncrementHitsCounter")
}


