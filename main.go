package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yoohahn/go-templ/handler"
	my_middleware "github.com/yoohahn/go-templ/middleware"
)

func setupMiddlewares(app *echo.Echo) {
	app.Pre(my_middleware.Auth())
	// app.Use(middleware.Logger())
	app.Use(middleware.Secure())
	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Browse: false,
		Root:   ".",
	}))
}

func setupRoot(app *echo.Echo) {
	rootHandler := handler.RootHandler{}
	app.GET("/", rootHandler.HandleRootShow)
}

func setupAuth(app *echo.Echo) {
	authHandler := handler.AuthHandler{}
	app.GET("/auth", authHandler.HandleAuthView)
}

func main() {
	app := echo.New()

	setupMiddlewares(app)
	// /
	setupRoot(app)
	// /auth
	setupAuth(app)

	fmt.Println("Server http://localhost:3000")
	app.Start(":3000")
}
