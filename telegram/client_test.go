package telegram

import (
	"testing"

	"github.com/mymmrac/telego"
)

func TestTgClient_processUpdate_faultyUpdate(t *testing.T) {
	tgClient := &TgClient{}

	update1 := telego.Update{}
	err := tgClient.ProcessUpdate(update1)
	if _, ok := err.(UnprocessableMessageError); !ok {
		t.Errorf("Expected UnprocessableMessageError, but got: %v", err)
	}
}

func TestTgClient_processUpdate_startCommand(t *testing.T) {
	tgClient := &TgClient{
		botApi: &telego.Bot{},
	}

	update := telego.Update{
		Message: &telego.Message{
			Text: "/start",
			Chat: telego.Chat{},
		},
	}
	err := tgClient.ProcessUpdate(update)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestTgClient_processUpdate_missingCommand(t *testing.T) {
	tgClient := &TgClient{}

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
	tgClient := &TgClient{}

	updateText := "Hello, world!"
	update := telego.Update{
		Message: &telego.Message{
			Text: updateText,
			Chat: telego.Chat{},
		},
	}
	err := tgClient.ProcessUpdate(update)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
