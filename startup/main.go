package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Anatolij-Grigorjev/tele-go-chi/interactions"
	"github.com/Anatolij-Grigorjev/tele-go-chi/pets_handling"
	"github.com/Anatolij-Grigorjev/tele-go-chi/storage"
	"github.com/Anatolij-Grigorjev/tele-go-chi/telegram"
	"github.com/mymmrac/telego"
)

var petsRepository storage.PetsRepository
var petsService *pets_handling.PetsService

func main() {

	prepareDataStore()
	petsRepository, _ = storage.NewDBPetsRepository(nil)
	petsService, _ = pets_handling.NewPetsService(petsRepository)
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
			return printBotGreeting()
		},
		"newpet": func(tgUpdate telego.Update) (string, error) {
			return tryCreateNewPet(tgUpdate)
		},
	}
}

func printBotGreeting() (string, error) {
	return interactions.START_GREETING, nil
}

func tryCreateNewPet(tgUpdate telego.Update) (string, error) {
	playerId := strconv.FormatInt(tgUpdate.Message.From.ID, 10)
	cmdArgs := strings.Split(tgUpdate.Message.Text, " ")
	var petEmoji string
	if len(cmdArgs) < 2 {
		petEmoji = ""
	} else {
		petEmoji = cmdArgs[1]
	}
	newPet, err := petsService.StoreNewPlayerPet(playerId, petEmoji)
	if err != nil {
		return err.Error(), nil
	}
	return interactions.NewPetMessage(newPet), nil
}
