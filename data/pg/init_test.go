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
	conn, err := model.Open("postgres", "postgres://xgmethyc:h2KYmnYJ15ZezhXWOB5NzwFBCNK55P7D@stampy.db.elephantsql.com:5432/xgmethyc") // local/test instance
	if err != nil {
		fmt.Println(err)
	}
	db = conn
	defer conn.Close()

	if err = engine.ResetTestMigrations(); err != nil {
		fmt.Println(err)
	}

	os.Exit(m.Run())
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
	return
}
