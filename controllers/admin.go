package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/rgomezs4/event_registration/data"
	"github.com/rgomezs4/event_registration/data/model"
	"github.com/rgomezs4/event_registration/engine"
)

// Admin handles every request /admin/xxx
type Admin struct {
}

func newAdmin() *engine.Route {
	var a interface{} = Admin{}
	return &engine.Route{
		Logger:  true,
		Handler: a.(http.Handler),
	}
}

func (ad Admin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "" && r.Method == "POST" {
		ad.create(w, r)
		return
	} else if head == "" && r.Method == "GET" {
		ad.find(w, r)
		return
	} else if head == "" && r.Method == "PUT" {
		ad.update(w, r)
		return
	} else if head == "login" && r.Method == "POST" {
		ad.login(w, r)
		return
	} else if head == "password" && r.Method == "PUT" {
		ad.updatePassword(w, r)
		return
	}

	newError(fmt.Errorf("path not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
}

func (ad Admin) create(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	// Converts the payload to a FinanceBatch object
	var data model.JSONApiRequest
	var admin model.Admin
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &admin)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	admin.ID, err = db.Admin.Insert(tx, admin)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}
	_ = engine.Respond(w, r, http.StatusCreated, admin)
}

func (ad Admin) find(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	adminID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newError(errors.New("invalid AdminID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	admin, err := db.Admin.Find(tx, adminID)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if admin == nil {
		_ = tx.Commit()
		newError(errors.New("admin not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}
	_ = engine.Respond(w, r, http.StatusOK, admin)
}

func (ad Admin) update(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	adminID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newError(errors.New("invalid AdminID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}
	// Converts the payload to a FinanceBatch object
	var data model.JSONApiRequest
	var admin model.UpdateAdmin
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &admin)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	a, err := db.Admin.Update(tx, adminID, admin)
	switch {
	case err == sql.ErrNoRows:
		_ = tx.Rollback()
		newError(fmt.Errorf("User with id %d not found", adminID), http.StatusNotFound).Handler.ServeHTTP(w, r)
		return
	case err != nil:
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}
	_ = engine.Respond(w, r, http.StatusOK, a)
}

func (ad Admin) login(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	// Converts the payload to a FinanceBatch object
	var data model.JSONApiRequest
	var log model.Login
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &log)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	a, err := db.Admin.Login(tx, log.Username, log.Password)
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if a == nil {
		newError(errors.New("incorrect username or password"), http.StatusNotFound).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	ctx = context.WithValue(ctx, engine.ContextAuth, a.ID)

	_ = engine.Respond(w, r.WithContext(ctx), http.StatusOK, a)
}

func (ad Admin) updatePassword(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newError(errors.New("invalid ID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}
	// Converts the payload to a FinanceBatch object
	var data model.JSONApiRequest
	var login model.Login
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &login)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := db.Admin.UpdatePassword(tx, id, login.Password); err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}
	var response = make(map[string]string)
	response["message"] = "Password updated successfully"
	_ = engine.Respond(w, r, http.StatusOK, response)
}
