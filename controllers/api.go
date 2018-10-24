package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/rgomezs4/event_registration/data"
	"github.com/rgomezs4/event_registration/engine"
)

// API is the starting point of our API.
// Responsible for routing the request to the correct handler
type API struct {
	DB            *data.DB
	Logger        func(http.Handler) http.Handler
	Authenticator func(http.Handler) http.Handler
}

// NewAPI returns a production API with all middlewares
func NewAPI() *API {
	return &API{
		Logger:        engine.Logger,
		Authenticator: engine.AuthenticationHandler,
	}
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, engine.ContextOriginalPath, r.URL.Path)
	ctx = context.WithValue(ctx, engine.ContextDatabase, a.DB)

	var next *engine.Route
	var head string

	uID := r.Header.Get("AUTH_USER_ID")
	if uID != "" {
		userID, err := strconv.Atoi(uID)
		if err != nil {
			next = newError(fmt.Errorf("path not found"), http.StatusNotFound)
		} else {
			ctx = context.WithValue(ctx, engine.ContextAuth, userID)
		}
	}

	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "admin" {
		next = newAdmin()
	} else if head == "person" {
		next = newPerson()
	} else if head == "item" {
		next = newItem()
	} else if head == "product" {
		next = newProduct()
	} else if head == "order" {
		next = newOrder()
	} else if head == "health" {
		next = health()
	} else {
		next = newError(fmt.Errorf("path not found"), http.StatusNotFound)
	}

	next.Handler = a.Authenticator(next.Handler)

	if next.Logger {
		next.Handler = a.Logger(next.Handler)
	}

	next.Handler.ServeHTTP(w, r.WithContext(ctx))
}

func health() *engine.Route {
	res := make(map[string]interface{})
	res["is_up"] = true

	return &engine.Route{
		Logger: true,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			engine.Respond(w, r, http.StatusOK, res)
		}),
	}
}

func newError(err error, statusCode int) *engine.Route {
	return &engine.Route{
		Logger: true,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			engine.Respond(w, r, statusCode, err)
		}),
	}
}

func saveFile(w http.ResponseWriter, r *http.Request, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
	if err != nil {
		engine.Respond(w, r, http.StatusBadRequest, err)
		return
	}

	engine.Respond(w, r, http.StatusOK, handle.Filename)
}
