package telegram

import (
	"testing"

	"github.com/mymmrac/telego"
	"go.uber.org/mock/gomock"
)

func conditionSendParamsHaveText(text string) gomock.Matcher {
	return gomock.Cond(func(x any) bool { return x.(*telego.SendMessageParams).Text == text })
}

func NewTgClientWithMockApi(t *testing.T) (*TgClient, *MockTelegoBotApi) {
	ctrl := gomock.NewController(t)
	botApi := NewMockTelegoBotApi(ctrl)
	return &TgClient{botApi: botApi}, botApi
}

func TestTgClient_processUpdate_faultyUpdate(t *testing.T) {
	tgClient, _ := NewTgClientWithMockApi(t)

	update1 := telego.Update{}
	err := tgClient.ProcessUpdate(update1)
	if _, ok := err.(UnprocessableMessageError); !ok {
		t.Errorf("Expected UnprocessableMessageError, but got: %v", err)
	}
}

func TestTgClient_processUpdate_startCommand(t *testing.T) {
	tgClient, botApi := NewTgClientWithMockApi(t)

	update := telego.Update{
		Message: &telego.Message{
			Text: "/start",
			Chat: telego.Chat{},
		},
	}

	botApi.EXPECT().SendMessage(conditionSendParamsHaveText(_START_GREETING)).Times(1)
	err := tgClient.ProcessUpdate(update)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestTgClient_processUpdate_missingCommand(t *testing.T) {
	tgClient, _ := NewTgClientWithMockApi(t)

	cmdWord := "potato"
	update := telego.Update{
		Message: &telego.Message{
			Text: "/" + cmdWord,
			Chat: telego.Chat{},
		},
	}
	err := tgClient.ProcessUpdate(update)
	missingCmdErr, typesMatch := err.(MissingCommandError)
	if !typesMatch {
		t.Errorf("Expected MissingCommandError, but got: %v", err)
	}
	if missingCmdErr.Command != cmdWord {
		t.Errorf("Expected command: %v, but got: %v", cmdWord, missingCmdErr.Command)
	}
}

func TestTgClient_processUpdate_echoMessage(t *testing.T) {
	tgClient, botApi := NewTgClientWithMockApi(t)

	updateText := "Hello, world!"
	update := telego.Update{
		Message: &telego.Message{
			Text: updateText,
			Chat: telego.Chat{},
		},
	}
	botApi.EXPECT().SendMessage(conditionSendParamsHaveText(updateText)).Times(1)
	err := tgClient.ProcessUpdate(update)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
