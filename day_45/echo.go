package day_45

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func Echo() {
	r := echo.New()

	// Query string
	// http://localhost:3000/page1?name=grayson
	r.GET("/profile", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("What's up %s", name)

		return ctx.String(http.StatusOK, data)
	})

	// http://localhost:3000/page2/grayson
	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	// Form value
	r.POST("/register", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		data := fmt.Sprintf(
			"Hello %s, I have message for you: %s",
			name,
			strings.Replace(message, "/", "", 1),
		)

		return ctx.String(http.StatusOK, data)
	})

	r.Start("127.0.0.1:3000")
}
