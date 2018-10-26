package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/rgomezs4/event_registration/data/model"
)

func TestItem_create(t *testing.T) {
	payload := []byte(`{"data":{"type":"create","attributes":{"name":"Test Item"}}}`)

	req, err := http.NewRequest("POST", "/item", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	item := &model.Item{}
	if status := rec.Code; status != http.StatusCreated {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusCreated)
	}

	if err := json.Unmarshal(body, item); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if (&model.Item{}) == item {
		t.Fatal("Item is empty")
	}
}

func TestItem_all(t *testing.T) {
	req, err := http.NewRequest("GET", "/item", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	items := &[]model.Item{}
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
