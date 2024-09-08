package pets_handling

import (
	"errors"
	"strings"

	"github.com/Anatolij-Grigorjev/tele-go-chi/storage"
	"github.com/google/uuid"
)

type PetsService struct {
	petsRepo storage.PetsRepository
}

func NewPetsService(petsRepo storage.PetsRepository) (*PetsService, error) {
	return &PetsService{petsRepo: petsRepo}, nil
}

func (service *PetsService) StoreNewPlayerPet(playerId string, petEmoji string) (storage.PlayerPet, error) {
	if strings.TrimSpace(playerId) == "" {
		return storage.PlayerPet{}, errors.New("missing meaningful player ID")
	}
	if strings.TrimSpace(petEmoji) == "" {
		return storage.PlayerPet{}, errors.New("missing pet emoji")
	}

	pet, err := createPlayerPet(playerId, petEmoji)
	if err != nil {
		return storage.PlayerPet{}, err
	}

	return service.petsRepo.SavePet(pet)
}

func createPlayerPet(playerId string, emoji string) (storage.PlayerPet, error) {
	petUUID, err := uuid.NewRandom()
	if err != nil {
		return storage.PlayerPet{}, err
	}

	return storage.PlayerPet{
		PlayerID: playerId,
		PetUUID:  petUUID.String(),
		PetEmoji: emoji,
	}, nil
}
