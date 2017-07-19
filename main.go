package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"net/http"
	"os"
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

	http.HandleFunc("/callback", func(w http.RequestWriter, req *http.Request) {
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
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						panic(err)
					}
				}
			}
		}
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
