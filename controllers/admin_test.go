package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/rgomezs4/event_registration/data/model"
)

func TestAdmin_create(t *testing.T) {
	payload := []byte(`{"data":{"type":"create","attributes":{"username":"user1","password":"123456","name":"TUser"}}}`)

	req, err := http.NewRequest("POST", "/admin", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	admin := &model.Admin{}
	if status := rec.Code; status != http.StatusCreated {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusCreated)
	}

	if err := json.Unmarshal(body, admin); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if (&model.Admin{}) == admin {
		t.Fatal("Admin is empty")
	}
}

func TestAdmin_find(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin?id=1", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	admin := &model.Admin{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusCreated)
	}

	if err := json.Unmarshal(body, admin); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if admin.Name == "" {
		t.Errorf("Returns %v expecting %v", "", "something")
	}
}

func TestAdmin_update(t *testing.T) {
	payload := []byte(`{"data":{"type":"update","attributes":{"name":"TUser1"}}}`)

	req, err := http.NewRequest("PUT", "/admin?id=1", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	admin := &model.Admin{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}

	if err := json.Unmarshal(body, admin); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if (&model.Admin{}) == admin {
		t.Fatal("Admin is empty")
	}
}

func TestAdmin_login(t *testing.T) {
	payload := []byte(`{"data":{"type":"login","attributes":{"username":"user1","password":"123456"}}}`)

	req, err := http.NewRequest("POST", "/admin/login", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	admin := &model.Admin{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}

	if err := json.Unmarshal(body, admin); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if (&model.Admin{}) == admin {
		t.Fatal("Admin	 is empty")
	}
}

func TestAdmin_updatePassword(t *testing.T) {
	payload := []byte(`{"data":{"type":"login","attributes":{"username":"user1","password":"654321"}}}`)

	req, err := http.NewRequest("PUT", "/admin/password?id=1", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	checkError(t, err)

	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}
}
