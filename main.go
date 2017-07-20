package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func envLoad() error {
	return godotenv.Load()
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())

	e.Static("/assets", "assets")

	if err := envLoad(); err != nil {
		panic(err)
	}

	handle(e)

	e.Logger.Fatal(e.Start(":8081"))
}
