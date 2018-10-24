package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rgomezs4/event_registration/data"
	"github.com/rgomezs4/event_registration/data/model"
	"github.com/rgomezs4/event_registration/engine"
)

// Item handles every request /item/xxx
type Item struct {
}

func newItem() *engine.Route {
	var i interface{} = Item{}
	return &engine.Route{
		Logger:  true,
		Handler: i.(http.Handler),
	}
}

func (it Item) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "" && r.Method == "POST" {
		it.create(w, r)
		return
	} else if head == "" && r.Method == "GET" {
		it.all(w, r)
		return
	}

	newError(fmt.Errorf("path not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
}

func (it Item) create(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	// Converts the payload to a FinanceBatch object
	var data model.JSONApiRequest
	var item model.Item
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &item)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	item.ID, err = db.Item.InsertItem(tx, item)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}
	_ = engine.Respond(w, r, http.StatusCreated, item)
}

func (it Item) all(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	items, err := db.Item.AllItems(tx)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	_ = engine.Respond(w, r, http.StatusOK, items)
}
