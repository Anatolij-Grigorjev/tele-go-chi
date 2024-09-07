package main

import (
	"fmt"
	"os"

	"github.com/Anatolij-Grigorjev/tele-go-chi/interactions"
	"github.com/Anatolij-Grigorjev/tele-go-chi/storage"
	"github.com/Anatolij-Grigorjev/tele-go-chi/telegram"
	"github.com/mymmrac/telego"
)

func main() {

	prepareDataStore()
	processTelegramUpdates()
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func prepareDataStore() {
	dbCredentials := storage.Credentials{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	err := storage.RunMigrations(dbCredentials)
	exitOnError(err)
}

func processTelegramUpdates() {
	tgBotToken := os.Getenv("BOT_TOKEN")

	botClient, err := telegram.NewTgClient(tgBotToken, buildTgHandlers())
	exitOnError(err)

	tgUpdates, err := botClient.OpenUpdatesChannel()
	exitOnError(err)

	for update := range tgUpdates {
		err := botClient.ProcessUpdate(update)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func buildTgHandlers() map[string]telegram.TgUpdateHandler {
	return map[string]telegram.TgUpdateHandler{
		"start": func(tgUpdate telego.Update) (string, error) {
			return interactions.START_GREETING, nil
		},
	}
}
