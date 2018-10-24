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
	m, err := migrate.New("file:///"+basepath, "postgres://xgmethyc:h2KYmnYJ15ZezhXWOB5NzwFBCNK55P7D@stampy.db.elephantsql.com:5432/xgmethyc")

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
