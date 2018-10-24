// +build !mem

package model

import (
	"database/sql"
	// Sets the database driver to postgres
	_ "github.com/lib/pq"
)

// Connection - Defines the connection type to be a sql connection
type Connection = sql.DB

// Key - defines the key type for pg models
type Key = int

// Open - opens the connection with the pg database and returns the sql.DB object
func Open(options ...string) (*sql.DB, error) {
	conn, err := sql.Open(options[0], options[1])
	if err != nil {
		return nil, err
	} else if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
