package main

import (
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"net/http"
	"os"
	"fmt"
)

func envLoad() error {
	return godotenv.Load()
}

func main() {

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

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("regist request.")

		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}

			return
		}

		for _, event := range events {
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
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("test")
        })

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
