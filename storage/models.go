package storage

import (
	"time"

	"github.com/upper/db/v4"
)

type PlayerPet struct {
	ID        int64     `db:"id,omitempty"`
	PlayerID  string    `db:"player_id"`
	PetUUID   string    `db:"pet_uuid"`
	PetEmoji  string    `db:"pet_emoji"`
	Alive     bool      `db:"pet_alive,omitempty"`
	CreatedAt time.Time `db:"created_at,omitempty"`
}

func (p *PlayerPet) Store(session db.Session) db.Store {
	return session.Collection("player_pets")
}
