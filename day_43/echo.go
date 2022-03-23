package day_43

import (
    "github.com/labstack/echo"
    "net/http"
)

type M map[string]interface{}

func Echo(){
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Echo says hello!"
		return ctx.String(http.StatusOK, data)
	})

	r.Start("127.0.0.1:9000")
}