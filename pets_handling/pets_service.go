package pets_handling

import "github.com/Anatolij-Grigorjev/tele-go-chi/storage"

type PetsService struct {
	repo storage.PetsRepository
}
