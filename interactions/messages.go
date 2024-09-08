package interactions

import (
	"fmt"

	"github.com/Anatolij-Grigorjev/tele-go-chi/storage"
)

const START_GREETING = `
[WIP]🚧🚧🚧🚧🚧🚧🚧[WIP]
Hello! 
Welcome to Tele-go-chi bot!
Choose your desired animal: 🐕 🐈 🦎 🦆
Watch them evolve and grow as they
* eat when they are hungry
* play when they are lonely
* poop after they eat
etc...

Use the ` + "`/newpet <emoji>`" + ` command to get started!
`

const _NEW_PET_TEMPLATE = `
New pet registered! 🎉
Say hello to your new pet %s:
` + "```" + `
%s
` + "```" + `
has a nice ring to it! 😊
`

func NewPetMessage(pet storage.PlayerPet) string {
	return fmt.Sprintf(_NEW_PET_TEMPLATE, pet.PetEmoji, pet.PetUUID)
}
