package main

import (
	// "netl/pkg/db"
	handler "netl/pkg/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	h := handler.Handler{}
	e.POST("/register/", h.RegisterUser)
	e.POST("/login/", h.LoginUser)

	e.Logger.Fatal(e.Start(":3000"))
}
