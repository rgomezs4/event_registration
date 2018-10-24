package controllers

import (
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

// Product handles every request /product/xxx
type Product struct {
}

func newProduct() *engine.Route {
	var p interface{} = Product{}
	return &engine.Route{
		Logger:  true,
		Handler: p.(http.Handler),
	}
}

func (pr Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "" && r.Method == "POST" {
		pr.create(w, r)
		return
	} else if head == "" && r.Method == "GET" {
		if id := r.URL.Query().Get("id"); id != "" {
			pr.find(w, r)
		} else if barcode := r.URL.Query().Get("barcode"); barcode != "" {
			pr.findByBarcode(w, r)
		} else {
			pr.all(w, r)
		}
		return
	} else if head == "" && r.Method == "PUT" {
		pr.update(w, r)
		return
	}

	newError(fmt.Errorf("path not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
}

func (pr Product) create(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	// Converts the payload to an object
	var data model.JSONApiRequest
	var product model.Product
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &product)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	product.ID, err = db.Product.Insert(tx, product)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}
	_ = engine.Respond(w, r, http.StatusCreated, product)
}

func (pr Product) all(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	products, err := db.Product.All(tx)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if products == nil {
		products = make([]model.Product, 0)
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	_ = engine.Respond(w, r, http.StatusOK, products)
}

func (pr Product) find(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	productID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newError(errors.New("invalid productID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	product, err := db.Product.Find(tx, productID)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if product == nil {
		_ = tx.Commit()
		newError(errors.New("product not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	_ = engine.Respond(w, r, http.StatusOK, product)
}

func (pr Product) findByBarcode(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	barcode := r.URL.Query().Get("barcode")

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	product, err := db.Product.FindByBarcode(tx, barcode)
	if err != nil {
		_ = tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if product == nil {
		_ = tx.Commit()
		newError(errors.New("product not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	_ = engine.Respond(w, r, http.StatusOK, product)
}

func (pr Product) update(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	productID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newError(errors.New("invalid productID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}
	// Converts the payload to a FinanceBatch object
	var data model.JSONApiRequest
	var product model.Product
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &product)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	p, err := db.Product.Update(tx, productID, product)
	switch {
	case err == sql.ErrNoRows:
		_ = tx.Rollback()
		newError(fmt.Errorf("User with id %d not found", productID), http.StatusNotFound).Handler.ServeHTTP(w, r)
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
	_ = engine.Respond(w, r, http.StatusOK, p)
}
