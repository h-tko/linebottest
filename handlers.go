package main

import (
	"github.com/labstack/echo"
)

func handle(e *echo.Echo) {
	e.POST("/callback/", func(c echo.Context) error {
		fmt.Println("regist request.")

		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			return err
		}

		for _, event := range events {
			// userID取り出す
			userID = event.Source.UserID

			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					var send string
					switch message.Text {
					case "あ":
						send = "ありがとう！"
					case "い":
						send = "まじかよ"
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(send)).Do(); err != nil {
						fmt.Printf("%v", err)
					}
				}
			}
		}

		fmt.Printf("userid: %s\n", userID)

		return nil
	})

	// pushAPI
	e.GET("/push/", func(c echo.Context) error {
		fmt.Println("regist push.")

		if _, err := bot.PushMessage(userID, linebot.NewTextMessage("push")).Do(); err != nil {
			fmt.Printf("%v", err)
		}

		return nil
	})
}
