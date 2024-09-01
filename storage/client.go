package storage

import (
	db "github.com/upper/db/v4"
)

type DBClient struct {
	dbSession db.Session
	petsStore PetsRepository
}
