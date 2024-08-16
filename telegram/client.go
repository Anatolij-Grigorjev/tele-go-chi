package telegram

import (
	"fmt"

	"github.com/Anatolij-Grigorjev/tele-go-chi/utils"
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
	defer utils.SetUpProcessInterrupt(client.StopUpdates)
	return client, nil
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
	tgClient.botApi.StopLongPolling()
	fmt.Println("\nStopping bot, bye-bye!")
}

func (tgClient *TgClient) ProcessMessage(update telego.Update) error {
	if cannotProcessUpdate(update) {
		return UnprocessableMessageError{}
	}

	return tgClient.echoMessage(update.Message)
}

func cannotProcessUpdate(update telego.Update) bool {
	return update.Message == nil && update.CallbackQuery == nil
}

func (tgClient *TgClient) echoMessage(message *telego.Message) error {
	chatId := message.Chat.ChatID()

	_, err := tgClient.botApi.CopyMessage(&telego.CopyMessageParams{
		ChatID:     chatId,
		FromChatID: chatId,
		MessageID:  message.MessageID,
	})

	return err
}
