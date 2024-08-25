package storage

import (
	"os"

	"github.com/hashicorp/go-set/v2"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

type Credentials struct {
	Username, Password string
}

const _MIGRATIONS_FOLDER_PATH = "db_migrations/"
const _MIGRATIONS_TABLE_NAME = "schema_migrations"
const _MIGRATIONS_TABLE_DDL = `
CREATE TABLE IF NOT EXISTS schema_migrations (
	version VARCHAR(255) NOT NULL PRIMARY KEY
);
`

func RunPendingMigrations(credentials Credentials) error {

	session, err := openSession(credentials)
	if err != nil {
		return err
	}
	defer session.Close()

	err = assertMigrationsLogExists(session)
	if err != nil {
		return err
	}

	filenames, err := fetchPossibleMigrationFiles(_MIGRATIONS_FOLDER_PATH)
	if err != nil {
		return err
	}

	loggedMigrations, err := fetchLoggedMigrations(session)
	if err != nil {
		return err
	}

	err = runMissingMigrations(session, filenames, loggedMigrations)
	return err
}

func openSession(credentials Credentials) (db.Session, error) {
	settings := mysql.ConnectionURL{
		Database: "telegochi",
		Host:     "localhost",
		User:     credentials.Username,
		Password: credentials.Password,
	}
	return mysql.Open(settings)
}

func assertMigrationsLogExists(session db.Session) error {
	ok, err := session.Collection("schema_migrations").Exists()
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	_, err = session.SQL().Exec(_MIGRATIONS_TABLE_DDL)
	return err
}

func fetchPossibleMigrationFiles(folderPath string) ([]string, error) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}
	var filenames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filenames = append(filenames, file.Name())
	}
	return filenames, nil
}

func fetchLoggedMigrations(session db.Session) (*set.Set[string], error) {
	var loggedMigrations []string
	err := session.Collection(_MIGRATIONS_TABLE_NAME).Find("version").OrderBy("version").All(&loggedMigrations)
	return set.From(loggedMigrations), err
}

func runMissingMigrations(session db.Session, possibleMigrationsFilenames []string, presentMigrations *set.Set[string]) error {
	for _, filename := range possibleMigrationsFilenames {
		if presentMigrations.Contains(filename) {
			continue
		}
		err := runMigration(session, filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func runMigration(session db.Session, filename string) error {
	migration, err := os.ReadFile(_MIGRATIONS_FOLDER_PATH + filename)
	if err != nil {
		return err
	}
	_, err = session.SQL().Exec(string(migration))
	if err != nil {
		return err
	}
	_, err = session.Collection(_MIGRATIONS_TABLE_NAME).Insert(map[string]string{"version": filename})
	return err
}
