package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"os"
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

	bot, err := linebot.New(
		os.Getenv("LINEBOT_SECRET"),
		os.Getenv("LINEBOT_TOKEN"),
	)

	if err != nil {
		panic(err)
	}

	// userIDを保存しておく場所（お試しなのでメモリで持っとく）
	var userID string

	handle(e)

	e.Logger.Fatal(e.Start(":8081"))
}
