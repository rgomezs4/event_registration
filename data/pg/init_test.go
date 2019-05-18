package pg

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/rgomezs4/event_registration/data/model"
	"github.com/rgomezs4/event_registration/engine"
)

var db *sql.DB

func TestMain(m *testing.M) {
	dbSource := os.Getenv("APP_DB_SOURCE_M")
	dbDriver := os.Getenv("APP_DB_DRIVER")
	conn, err := model.Open(dbDriver, dbSource)
	if err != nil {
		fmt.Println(err)
	}
	db = conn
	defer conn.Close()

	if err = engine.ResetTestMigrations("data"); err != nil {
		fmt.Println(err)
	}

	os.Exit(m.Run())
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
