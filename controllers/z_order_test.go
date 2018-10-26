package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/rgomezs4/event_registration/data/model"
)

func TestOrder_create(t *testing.T) {
	payload := []byte(`{"data":{"type":"login","user_id":1,"attributes":{"person_id":1,"payment_method":1,"total":150,"comment":"none","btc_address":"address","detail":[{"product_id":1,"quantity":10,"price":15,"amount":150}]}}}`)

	req, err := http.NewRequest("POST", "/order", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	order := &model.OrderHeader{}
	if status := rec.Code; status != http.StatusCreated {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusCreated)
	}

	if err := json.Unmarshal(body, order); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if (&model.OrderHeader{}) == order {
		t.Fatal("Item is empty")
	}
}

func TestOrder_find(t *testing.T) {
	req, err := http.NewRequest("GET", "/order?id=1", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	order := &model.OrderHeader{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}

	if err := json.Unmarshal(body, order); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if order.Comment == "" {
		t.Errorf("Returns %v expecting %v", "", "something")
	}
}

func TestOrder_userOrders(t *testing.T) {
	req, err := http.NewRequest("GET", "/order?user_id=1", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	order := &[]model.OrderHeader{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}

	if err := json.Unmarshal(body, order); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if len(*order) == 0 {
		t.Errorf("Returns %d expecting >= %d", len(*order), 1)
	}
}

func TestOrder_totals(t *testing.T) {
	req, err := http.NewRequest("GET", "/order/totals?user_id=1", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	checkError(t, err)

	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}
}
