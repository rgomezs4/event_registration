package engine

import (
	"fmt"
	"os"
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
		db = os.Getenv("APP_DB_SOURCE_C")
	} else {
		db = os.Getenv("APP_DB_SOURCE_M")
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
	err = m.Up()
	return err
}
