package telegram

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/mymmrac/telego"
)

type TgClient struct {
	botApi *telego.Bot
}

func NewTgClient(token string) (*TgClient, error) {
	botClient, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil {
		return nil, err
	}
	client := &TgClient{botApi: botClient}
	defer client.setUpInterrupt()
	return client, nil
}

func (tgClient *TgClient) setUpInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			tgClient.StopUpdates()
		}
	}()
}

func (tgClient *TgClient) OpenUpdatesChannel() (<-chan telego.Update, error) {
	longPollingParams := telego.GetUpdatesParams{
		Limit:   150,
		Timeout: 5,
	}

	tgUpdates, err := tgClient.botApi.UpdatesViaLongPolling(&longPollingParams)
	if err != nil {
		return nil, err
	}
	return tgUpdates, nil
}

func (tgClient *TgClient) StopUpdates() {
	fmt.Println("Stopping bot, bye-bye!")
	tgClient.botApi.StopLongPolling()
}

func (tgClient *TgClient) EchoMessage(message *telego.Message) error {
	chatId := message.Chat.ChatID()

	_, err := tgClient.botApi.CopyMessage(&telego.CopyMessageParams{
		ChatID:     chatId,
		FromChatID: chatId,
		MessageID:  message.MessageID,
	})

	return err
}
