package storage

import (
	"time"
)

type PlayerPet struct {
	ID        int64     `db:"id,omitempty"`
	PlayerID  string    `db:"player_id"`
	PetUUID   string    `db:"pet_uuid"`
	PetEmoji  string    `db:"pet_emoji"`
	Alive     bool      `db:"pet_alive,omitempty"`
	CreatedAt time.Time `db:"created_at,omitempty"`
}
