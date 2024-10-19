package storage

import (
	"database/sql"

	"github.com/pressly/goose/v3"
	"github.com/upper/db/v4"
)

const _MIGRATIONS_FOLDER_PATH = "storage/db_migrations/"

func RunMigrations(dbCredentials Credentials) error {
	return WithSession(dbCredentials, runGooseMigrations)
}

func runGooseMigrations(session db.Session) error {
	sqlDB := session.Driver().(*sql.DB)

	err := goose.SetDialect("mysql")
	if err != nil {
		return err
	}

	err = goose.Up(sqlDB, _MIGRATIONS_FOLDER_PATH)
	if err != nil {
		return err
	}

	return nil
}
