package engine

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/golang-migrate/migrate"
	// postgres migrations driver
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	// _ "github.com/lib/pq"
)

// ResetTestMigrations reverts and runs every migration on /migrations folder
func ResetTestMigrations(source string) error {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	basepath = strings.Replace(basepath, "/engine", "/migrations", -1)

	db := ""
	if source == "controllers" {
		db = "postgres://postgres:abc123@142.93.56.8:5432/c_test?sslmode=disable"
	} else {
		db = "postgres://postgres:abc123@142.93.56.8:5432/d_test?sslmode=disable"
	}

	m, err := migrate.New("file:///"+basepath, db)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer m.Close()

	if err := m.Drop(); err != nil {
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil {
		return err
	}

	return nil
}
