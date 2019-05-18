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
	if err := engine.ResetTestMigrations("controllers"); err != nil {
		fmt.Println(err, " migration error")
	}

	os.Exit(m.Run())
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	dbSource := os.Getenv("APP_DB_SOURCE_C")
	dbDriver := os.Getenv("APP_DB_DRIVER")
	conn, err := model.Open(dbDriver, dbSource)
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

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
