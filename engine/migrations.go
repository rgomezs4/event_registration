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
func ResetTestMigrations() error {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	basepath = strings.Replace(basepath, "/engine", "/migrations", -1)
	m, err := migrate.New("file:///"+basepath, "postgres://postgres:abc123@142.93.56.8:5432/events_test?sslmode=disable")

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer m.Close()

	if err := m.Drop(); err != nil {
		return err
	}
	// Migrate all the way up ...
	if err := m.Up(); err != nil {
		return err
	}

	return nil
}
