package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/rgomezs4/event_registration/data"
	"github.com/rgomezs4/event_registration/data/model"
	"github.com/rgomezs4/event_registration/engine"
)

func TestMain(m *testing.M) {
	if err := engine.ResetTestMigrations(); err != nil {
		fmt.Println(err, " migration error")
	}

	os.Exit(m.Run())
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	conn, err := model.Open("postgres", "postgres://xgmethyc:h2KYmnYJ15ZezhXWOB5NzwFBCNK55P7D@stampy.db.elephantsql.com:5432/xgmethyc")
	if err != nil {
		log.Fatal("error while connecting to db ", err)
	}

	db := &data.DB{Connection: conn}

	api := &API{
		DB:            db,
		Logger:        logger,
		Authenticator: authenticator,
	}

	rec := httptest.NewRecorder()

	api.ServeHTTP(rec, req)
	return rec
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func setTestEnviromentalVars() {
	os.Setenv("APP_PORT", "3158")
	os.Setenv("APP_NAME", "events")
	os.Setenv("APP_DB_DRIVER", "postgres")
	os.Setenv("APP_DB_SOURCE", "postgres://xgmethyc:h2KYmnYJ15ZezhXWOB5NzwFBCNK55P7D@stampy.db.elephantsql.com:5432/xgmethyc")
	os.Setenv("APP_KEY", "secret")
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
