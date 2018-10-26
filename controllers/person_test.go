package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/rgomezs4/event_registration/data/model"
)

func TestPerson_create(t *testing.T) {
	payload := []byte(`{"data":{"type":"login","attributes":{"first_name":"First","last_name":"Last","birthdate":"2009-11-10 23:00:00 +0000 UTC m=+0.000000001","passport_number":"#1","country_origin":"Guatemala","country_birth":"Guatemala","language":"spanish","gender":"male","transafer":"?","mastercouncil":"?","image":"?","status":1,"section":"?","position":"?","notes":"?","updated_by":1}}}`)

	req, err := http.NewRequest("POST", "/person", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	person := &model.Person{}
	if status := rec.Code; status != http.StatusCreated {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusCreated)
	}

	if err := json.Unmarshal(body, person); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if (&model.Person{}) == person {
		t.Fatal("Item is empty")
	}
}

func TestPerson_find(t *testing.T) {
	req, err := http.NewRequest("GET", "/person?id=1", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	person := &model.Person{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusCreated)
	}

	if err := json.Unmarshal(body, person); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}
}

func TestPerson_all(t *testing.T) {
	req, err := http.NewRequest("GET", "/person", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	persons := &[]model.Person{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusCreated)
	}

	if err := json.Unmarshal(body, persons); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if len(*persons) == 0 {
		t.Errorf("expected at least 1 item to return got %d", len(*persons))
	}
}

func TestPerson_update(t *testing.T) {
	payload := []byte(`{"data":{"type":"update","attributes":{"first_name":"Test","last_name":"123","birthdate":"2009-11-10 23:00:00 +0000 UTC m=+0.000000001","passport_number":"123","country_origin":"Guatemala","country_birth":"Guatemala","language":"Spanish","gender":"Male","transafer":"?","mastercouncil":"?","image":"none.jpg","status":1,"section":"1","position":"1","notes":"none","updated_by":1}}}`)

	req, err := http.NewRequest("PUT", "/person?id=1", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	person := &model.Person{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}

	if err := json.Unmarshal(body, person); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if (&model.Person{}) == person {
		t.Fatal("Item is empty")
	}
}

func TestPerson_register(t *testing.T) {
	payload := []byte(`{"data":{"type":"login","user_id":1,"attributes":{"id":1,"image":"asistee_two.png","items":[{"id":1}]}}}`)
	req, err := http.NewRequest("POST", "/person/register", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	checkError(t, err)

	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}
}

func TestPerson_getItems(t *testing.T) {
	req, err := http.NewRequest("GET", "/person/items?id=1", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	items := &[]model.PersonItem{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusCreated)
	}

	if err := json.Unmarshal(body, items); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if len(*items) == 0 {
		t.Errorf("expected at least 1 item to return got %d", len(*items))
	}
}
