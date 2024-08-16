package telegram

import (
	"fmt"
	"strings"

	"github.com/Anatolij-Grigorjev/tele-go-chi/utils"
	"github.com/mymmrac/telego"
)

type TgClient struct {
	botApi TelegoBotApi
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

func (tgClient *TgClient) ProcessUpdate(update telego.Update) error {
	if cannotProcessUpdate(update) {
		return UnprocessableMessageError{}
	}

	if messageIsCommand(update) {
		return tgClient.processCommand(update)
	}
	return tgClient.echoMessage(update.Message)
}

func cannotProcessUpdate(update telego.Update) bool {
	return update.Message == nil && update.CallbackQuery == nil
}

func messageIsCommand(update telego.Update) bool {
	return update.Message != nil && strings.HasPrefix(update.Message.Text, "/")
}

func (tgClient *TgClient) processCommand(update telego.Update) error {
	command, _ := strings.CutPrefix(update.Message.Text, "/")

	switch command {
	case "start":
		return tgClient.sendMessage(update.Message.Chat.ChatID(), _START_GREETING)
	default:
		return MissingCommandError{Command: command}
	}
}

func (tgClient *TgClient) echoMessage(message *telego.Message) error {
	chatId := message.Chat.ChatID()
	return tgClient.sendMessage(chatId, message.Text)
}

func (tgClient *TgClient) sendMessage(chatId telego.ChatID, messageText string) error {
	_, err := tgClient.botApi.SendMessage(&telego.SendMessageParams{
		ChatID: chatId,
		Text:   messageText,
	})

	return err
}
