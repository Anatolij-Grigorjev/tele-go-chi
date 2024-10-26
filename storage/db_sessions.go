package storage

import (
	"fmt"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

// Open db session with provided credentials.
//
// Closing is responsibility of caller
func OpenSession(credentials Credentials) (db.Session, func(), error) {
	settings := mysql.ConnectionURL{
		Database: credentials.DBName,
		Host:     credentials.Host,
		User:     credentials.Username,
		Password: credentials.Password,
		// required to avoid collation utf8mb3 errors in upperdb adapter
		Options: map[string]string{
			"charset":   "utf8mb4",
			"collation": "utf8mb4_unicode_ci",
		},
	}
	session, err := mysql.Open(settings)
	closerFunc := func() { fmt.Println("\nClosing DB Session..."); session.Close() }
	return session, closerFunc, err
}

// Run actions within a db session.
//
// Opening and closing happens automatically around actions.
func WithSession(credentials Credentials, actions func(session db.Session) error) error {
	session, closer, err := OpenSession(credentials)
	if err != nil {
		return err
	}
	defer closer()
	err = actions(session)
	return err
}
