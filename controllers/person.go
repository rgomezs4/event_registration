package controllers

import (
	"bufio"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/rgomezs4/event_registration/data"
	"github.com/rgomezs4/event_registration/data/model"
	"github.com/rgomezs4/event_registration/engine"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Person handles every request /person/xxx
type Person struct {
}

func newPerson() *engine.Route {
	var p interface{} = Person{}
	return &engine.Route{
		Logger:  true,
		Handler: p.(http.Handler),
	}
}

func (pe Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "" && r.Method == "POST" {
		pe.create(w, r)
		return
	} else if head == "" && r.Method == "GET" {
		if id := r.URL.Query().Get("id"); id != "" {
			pe.find(w, r)
		} else {
			pe.all(w, r)
		}
		return
	} else if head == "" && r.Method == "PUT" {
		pe.update(w, r)
		return
	} else if head == "upload" {
		pe.upload(w, r)
		return
	} else if head == "register" && r.Method == "POST" {
		pe.register(w, r)
		return
	} else if head == "items" && r.Method == "GET" {
		pe.getItems(w, r)
		return
	} else if head == "image" && r.Method == "GET" {
		pe.getImage(w, r)
		return
	}

	newError(fmt.Errorf("path not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
}

func (pe Person) create(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	// Converts the payload to a FinanceBatch object
	var data model.JSONApiRequest
	var person model.Person
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &person)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	person.ID, err = db.Person.Insert(tx, data.Data.UserID, person)
	if err != nil {
		tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}
	engine.Respond(w, r, http.StatusCreated, person)
}

func (pe Person) find(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	personID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newError(errors.New("invalid personID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	person, err := db.Person.Find(tx, personID)
	if err != nil {
		tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if person == nil {
		tx.Commit()
		newError(errors.New("person not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	engine.Respond(w, r, http.StatusOK, person)
}

func (pe Person) all(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	persons, err := db.Person.All(tx)
	if err != nil {
		tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	engine.Respond(w, r, http.StatusOK, persons)
}

func (pe Person) update(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	personID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newError(errors.New("invalid personID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}
	// Converts the payload to a FinanceBatch object
	var data model.JSONApiRequest
	var person model.Person
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &person)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	p, err := db.Person.Update(tx, personID, data.Data.UserID, person)
	switch {
	case err == sql.ErrNoRows:
		tx.Rollback()
		newError(fmt.Errorf("User with id %d not found", personID), http.StatusNotFound).Handler.ServeHTTP(w, r)
		return
	case err != nil:
		tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}
	engine.Respond(w, r, http.StatusOK, p)
}

func (pe Person) upload(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("file")

	filename := r.Header.Get("X_FILE_NAME")
	fileextension := strings.Split(handle.Filename, ".")[1]
	handle.Filename = filename + "." + fileextension

	if err != nil {
		engine.Respond(w, r, http.StatusBadRequest, "The format file is not valid.")
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		saveFile(w, r, file, handle)
	case "image/png":
		saveFile(w, r, file, handle)
	default:
		engine.Respond(w, r, http.StatusBadRequest, "The format file is not valid.")
	}

}

func (pe Person) register(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	// Converts the payload to an object
	var data model.JSONApiRequest
	var person model.RegisterPerson
	defer r.Body.Close()
	request, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(request, &data)

	raw, err := json.Marshal(data.Data.Attributes)
	if err != nil {
		newError(fmt.Errorf("failed to read payload"), http.StatusNotAcceptable).Handler.ServeHTTP(w, r)
		return
	}
	_ = json.Unmarshal(raw, &person)

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	err = db.Person.Register(tx, person.PersonID, data.Data.UserID, person.Image)
	switch {
	case err == sql.ErrNoRows:
		tx.Rollback()
		newError(fmt.Errorf("User with id %d not found", person.PersonID), http.StatusNotFound).Handler.ServeHTTP(w, r)
		return
	case err != nil:
		tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	for _, item := range person.Items {
		pi := model.PersonItem{
			CreatedBy: data.Data.UserID,
			ItemID:    item.ID,
			PersonID:  person.PersonID,
		}
		_, err = db.Item.InsertPersonItem(tx, pi, data.Data.UserID)
		if err != nil {
			tx.Rollback()
			newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}
	engine.Respond(w, r, http.StatusOK, "registered succesfully")
}

func (pe Person) getItems(w http.ResponseWriter, r *http.Request) {
	// Gets the database object and connection from the context of the request
	ctx := r.Context()
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	personID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newError(errors.New("invalid personID"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}

	// Starts the transaction
	tx, err := db.Connection.Begin()
	if err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	items, err := db.Item.PersonItems(tx, personID)
	if err != nil {
		tx.Rollback()
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	if items == nil {
		items = make([]model.PersonItem, 0)
	}

	if err := tx.Commit(); err != nil {
		newError(err, http.StatusInternalServerError).Handler.ServeHTTP(w, r)
		return
	}

	engine.Respond(w, r, http.StatusOK, items)
}

func (pe Person) getImage(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("image")

	if imageName == "" {
		newError(errors.New("must provide an image name"), http.StatusBadRequest).Handler.ServeHTTP(w, r)
		return
	}

	imgFile, err := os.Open("./files/" + imageName) // a QR code image

	if err != nil {
		engine.Respond(w, r, http.StatusOK, err)
		return
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// if you create a new image instead of loading from file, encode the image to buffer instead with png.Encode()

	// png.Encode(&buf, image)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	mresponse := make(map[string]string)
	mresponse["image"] = imgBase64Str

	engine.Respond(w, r, http.StatusOK, mresponse)
}
