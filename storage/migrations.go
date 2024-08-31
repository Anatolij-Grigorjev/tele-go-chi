package storage

import (
	"database/sql"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"

	"github.com/pressly/goose/v3"
)

const _MIGRATIONS_FOLDER_PATH = "db_migrations/"

func RunMigrations(dbCredentials Credentials) error {
	session, err := openSession(dbCredentials)
	if err != nil {
		return err
	}
	defer session.Close()
	sqlDB := session.Driver().(*sql.DB)

	err = goose.SetDialect("mysql")
	if err != nil {
		return err
	}

	err = goose.Up(sqlDB, _MIGRATIONS_FOLDER_PATH)
	if err != nil {
		return err
	}

	return nil
}

func openSession(credentials Credentials) (db.Session, error) {
	settings := mysql.ConnectionURL{
		Database: credentials.DBName,
		Host:     credentials.Host,
		User:     credentials.Username,
		Password: credentials.Password,
	}
	return mysql.Open(settings)
}
