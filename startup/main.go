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
	"github.com/Anatolij-Grigorjev/tele-go-chi/utils"
	"github.com/mymmrac/telego"
	"github.com/upper/db/v4"
)

var petsRepository storage.PetsRepository
var petsService *pets_handling.PetsService

func main() {
	// Run all exit functions when server quits
	utils.SetUpProcessInterrupt()

	botSession := prepareDataStore()
	petsRepository, _ = storage.NewDBPetsRepository(botSession)
	petsService, _ = pets_handling.NewPetsService(petsRepository)
	processTelegramUpdates()
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func prepareDataStore() db.Session {
	dbCredentials := storage.Credentials{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	err := storage.RunMigrations(dbCredentials)
	exitOnError(err)

	session, closer, err := storage.OpenSession(dbCredentials)
	exitOnError(err)
	// close db session when server halts
	utils.AddOnExitFunc(closer)
	return session
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
			return createNewPet(tgUpdate)
		},
	}
}

func printBotGreeting() (string, error) {
	return interactions.START_GREETING, nil
}

func createNewPet(tgUpdate telego.Update) (string, error) {
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
		return "", err
	}
	return interactions.NewPetMessage(newPet), nil
}
