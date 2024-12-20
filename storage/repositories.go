package storage

import (
	"errors"

	db "github.com/upper/db/v4"
)

type DBPetsRepository struct {
	dbSession db.Session
}

func NewDBPetsRepository(dbSession db.Session) (*DBPetsRepository, error) {
	return &DBPetsRepository{dbSession: dbSession}, nil
}

func (impl *DBPetsRepository) SavePet(pet PlayerPet) (PlayerPet, error) {
	err := impl.dbSession.Save(&pet)
	return pet, err
}

func (impl *DBPetsRepository) FindAllPlayerPets(PlayerID string) ([]PlayerPet, error) {
	return []PlayerPet{}, errors.New("TODO")
}
