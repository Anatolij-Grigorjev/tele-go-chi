package telegram

import (
	"github.com/mymmrac/telego"
)

type TelegoBotApi interface {
	UpdatesViaLongPolling(params *telego.GetUpdatesParams, options ...telego.LongPollingOption) (<-chan telego.Update, error)
	StopLongPolling()
	SendMessage(params *telego.SendMessageParams) (*telego.Message, error)
}
