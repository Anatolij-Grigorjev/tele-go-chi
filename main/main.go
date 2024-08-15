package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
)

func main() {

	tgBotToken := os.Getenv("BOT_TOKEN")

	botClient, err := telego.NewBot(tgBotToken, telego.WithDefaultDebugLogger())
	exitOnError(err)

	longPollingParams := telego.GetUpdatesParams{
		Limit:   150,
		Timeout: 5,
	}

	tgUpdates, err := botClient.UpdatesViaLongPolling(&longPollingParams)
	exitOnError(err)
	defer botClient.StopLongPolling()

	for update := range tgUpdates {
		if update.Message != nil {
			chatId := update.Message.Chat.ChatID()

			botClient.CopyMessage(&telego.CopyMessageParams{
				ChatID:     chatId,
				FromChatID: chatId,
				MessageID:  update.Message.MessageID,
			})
		} else {
			fmt.Println("Update came with no message")
		}
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
