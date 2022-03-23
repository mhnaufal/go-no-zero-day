package day_44

import (
    "github.com/labstack/echo"
    "net/http"
)

type M map[string]interface{}

func Echo(){
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		// render a plain text | text/plain
		data := "Echo says hello!"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/about", func(ctx echo.Context) error {
		// render a HTML | text/html
		data := "<h1>I rendered as HTML</h1>!"
		return ctx.HTML(http.StatusOK, data)
	})

	r.GET("/me", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/about")
	})

	r.GET("/login", func(ctx echo.Context) error {
		// render a JSON type | application/json
		data := M{"message": "I am a JSON", "status": "SUCCESS"}
		return ctx.JSON(http.StatusOK, data)
	})

	r.Start("127.0.0.1:9000")
}
