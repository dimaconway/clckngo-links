package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/labstack/echo/v4"

    "clckngo/links/models"
)

// Implements `ServerInterface`
type ApiImplementation struct{}

func (w *ApiImplementation) GetLinks(ctx echo.Context) error {
    return ctx.String(http.StatusOK, "GetLinks")
}

func (w *ApiImplementation) StoreLink(ctx echo.Context) error {
    var err error

    link := new(models.Link)

    err = ctx.Bind(link)

    if err != nil {
        return err
    }

    strLink, err := json.Marshal(link)
    responseStr := "StoreLink " + string(strLink)

    return ctx.String(http.StatusOK, responseStr)
}

func (w *ApiImplementation) DeleteLink(ctx echo.Context, hash string) error {
    return ctx.String(http.StatusOK, "DeleteLink")
}

func (w *ApiImplementation) GetLink(ctx echo.Context, hash string) error {
    return ctx.String(http.StatusOK, "GetLink")
}

func (w *ApiImplementation) IncrementHitsCounter(ctx echo.Context, hash string) error {
    return ctx.String(http.StatusOK, "IncrementHitsCounter")
}
