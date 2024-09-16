package pets_handling

import (
	"testing"
	"time"

	"github.com/Anatolij-Grigorjev/tele-go-chi/storage"
	"go.uber.org/mock/gomock"
)

func NewPetsServiceWithMockRepos(t *testing.T) (*PetsService, *storage.MockPetsRepository) {
	ctrl := gomock.NewController(t)
	petsRepo := storage.NewMockPetsRepository(ctrl)
	return &PetsService{petsRepo: petsRepo}, petsRepo
}

func TestPetsService_storeNewPet_missingPlayerid(t *testing.T) {
	service, petsRepo := NewPetsServiceWithMockRepos(t)

	petsRepo.EXPECT().SavePet(gomock.Any()).Times(0)
	_, err := service.StoreNewPlayerPet("", "🦆")
	if err == nil {
		t.Errorf("Empty player Id should cause error!")
	}
}

func TestPetsService_storeNewPet_missingEmoji(t *testing.T) {
	service, petsRepo := NewPetsServiceWithMockRepos(t)

	petsRepo.EXPECT().SavePet(gomock.Any()).Times(0)
	_, err := service.StoreNewPlayerPet("playerId", "")
	if err == nil {
		t.Errorf("Empty emoji should cause error!")
	}
}

func TestPetsService_storeNewPet_returnsStored(t *testing.T) {
	service, petsRepo := NewPetsServiceWithMockRepos(t)

	petsRepo.EXPECT().SavePet(gomock.Any()).DoAndReturn(func(pet storage.PlayerPet) (storage.PlayerPet, error) {
		return pet, nil
	})
	playerId := "playerId"
	emoji := "🦆"
	pet, err := service.StoreNewPlayerPet(playerId, emoji)
	if err != nil {
		t.Errorf("mock storing player returned error %s", err)
	}
	if !pet.Alive {
		t.Errorf("new pet is not alive!")
	}
	if (pet.CreatedAt == time.Time{}) {
		t.Errorf("new pet does not have creation timestamp!")
	}
	if pet.PlayerID != playerId {
		t.Errorf("new pet has wrong player id %s", pet.PlayerID)
	}
	if pet.PetEmoji != emoji {
		t.Errorf("new pet has wrong emoji %s", pet.PetEmoji)
	}
	if pet.PetUUID == "" {
		t.Errorf("new pet missing UUID!")
	}
}
