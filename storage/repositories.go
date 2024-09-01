package storage

type PetsRepository interface {
	SavePet(pet PlayerPet) (PlayerPet, error)
	FindAllPlayerPets(PlayerID string) ([]PlayerPet, error)
}
