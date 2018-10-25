package controllers

import (
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

// Order handles every /order/xxx request
type Order struct {
}

func newOrder() *engine.Route {
	var o interface{} = Order{}
	return &engine.Route{
		Logger:  true,
		Handler: o.(http.Handler),
	}
}

func (or Order) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "" && r.Method == "POST" {
		or.create(w, r)
		return
	} else if head == "" && r.Method == "GET" {
		if id := r.URL.Query().Get("id"); id != "" {
			or.find(w, r)
			return
		} else if id := r.URL.Query().Get("user_id"); id != "" {
			or.getByUserID(w, r)
			return
		}
	} else if head == "totals" && r.Method == "GET" {
		or.getTotals(w, r)
		return
	}

	newError(fmt.Errorf("path not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
}

func (or Order) create(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	// Converts the payload to a FinanceBatch object
	var data model.JSONApiRequest
	var order model.OrderHeader
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &order)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	order.ID, err = db.Order.InsertHeader(tx, data.Data.UserID, order)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	for i := range order.Detail {
		order.Detail[i].ID, err = db.Order.InsertDetail(tx, order.ID, order.Detail[i])
		if err != nil {
			_ = tx.Rollback()
			newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	_ = engine.Respond(w, r, http.StatusCreated, order)
}

func (or Order) find(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	orderID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newError(errors.New("invalid orderID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	order, err := db.Order.FindHeader(tx, orderID)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if order == nil {
		_ = tx.Rollback()
		newError(errors.New("order not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
		return
	}

	order.Detail, err = db.Order.FindDetail(tx, orderID)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	_ = engine.Respond(w, r, http.StatusOK, order)
}

func (or Order) getByUserID(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		newError(errors.New("invalid userID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	order, err := db.Order.GetOrdersByID(tx, userID)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if order == nil {
		order = make([]model.OrderHeader, 0)
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	_ = engine.Respond(w, r, http.StatusOK, order)
}

func (or Order) getTotals(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		newError(errors.New("invalid userID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	quantity, total, err := db.Order.GetTotalsByUser(tx, userID)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	mresp := make(map[string]interface{})
	mresp["quantity"] = quantity
	mresp["total"] = total

	_ = engine.Respond(w, r, http.StatusOK, mresp)
}
