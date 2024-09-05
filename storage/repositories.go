package storage

import (
	"errors"

	db "github.com/upper/db/v4"
)

type DBPetsRepository struct {
	dbSession db.Session
}

func (impl *DBPetsRepository) SavePet(pet PlayerPet) (PlayerPet, error) {
	return pet, errors.New("TODO")
}

func (impl *DBPetsRepository) FindAllPlayerPets(PlayerID string) ([]PlayerPet, error) {
	return []PlayerPet{}, errors.New("TODO")
}
