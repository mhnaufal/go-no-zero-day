package day_43

import (
    "fmt"
    "github.com/labstack/echo"
    "net/http"
    "strings"
)

type M map[string]Interface{}

func Echo(){
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Echo says hello!"
		return ctx.String(http.StatusOK, data)
	})

	r.Start("127.0.0.1:9000")
}