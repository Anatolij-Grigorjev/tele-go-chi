package storage

import (
	"time"

	"github.com/google/uuid"
)

type PlayerPet struct {
	ID        int64     `db:"id,omitempty"`
	PlayerID  string    `db:"player_id"`
	PetUUID   string    `db:"pet_uuid"`
	PetEmoji  string    `db:"pet_emoji"`
	Alive     bool      `db:"pet_alive,omitempty"`
	CreatedAt time.Time `db:"created_at,omitempty"`
}

func NewPet(PlayerID string, PetEmoji string) (PlayerPet, error) {
	petUUID, err := uuid.NewRandom()
	if err != nil {
		return PlayerPet{}, err
	}

	return PlayerPet{
		PlayerID: PlayerID,
		PetUUID:  petUUID.String(),
		PetEmoji: PetEmoji,
	}, nil
}
