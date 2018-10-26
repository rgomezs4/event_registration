package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/rgomezs4/event_registration/data/model"
)

func TestProduct_create(t *testing.T) {
	payload := []byte(`{"data":{"type":"login","attributes":{"name":"new item 1","barcode":"123456","size":"S","stock":100,"price":10.32}}}`)

	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	product := &model.Product{}
	if status := rec.Code; status != http.StatusCreated {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusCreated)
	}

	if err := json.Unmarshal(body, product); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if (&model.Product{}) == product {
		t.Fatal("Item is empty")
	}
}

func TestProduct_find(t *testing.T) {
	req, err := http.NewRequest("GET", "/product?id=1", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	product := &model.Product{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}

	if err := json.Unmarshal(body, product); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if product.Name == "" {
		t.Errorf("Returns %v expecting %v", "", "something")
	}
}

func TestProduct_findBarcode(t *testing.T) {
	req, err := http.NewRequest("GET", "/product?barcode=123456", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	product := &model.Product{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}

	if err := json.Unmarshal(body, product); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if product.Name == "" {
		t.Errorf("Returns %v expecting %v", "", "something")
	}
}

func TestProduct_all(t *testing.T) {
	req, err := http.NewRequest("GET", "/product", nil)
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	products := &[]model.Product{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusCreated)
	}

	if err := json.Unmarshal(body, products); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if len(*products) == 0 {
		t.Errorf("expected at least 1 item to return got %d", len(*products))
	}
}

func TestProduct_update(t *testing.T) {
	payload := []byte(`{"data":{"type":"login","attributes":{"name":"new item 2","barcode":"123456","size":"S","stock":100,"price":10.32}}}`)

	req, err := http.NewRequest("PUT", "/product?id=1", bytes.NewBuffer(payload))
	checkError(t, err)
	req.Header.Add("AUTH_USER_ID", "1")

	rec := executeRequest(req)
	body, err := ioutil.ReadAll(rec.Body)
	checkError(t, err)

	product := &model.Product{}
	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("returns status %v was expecting %v", status, http.StatusOK)
	}

	if err := json.Unmarshal(body, product); err != nil {
		t.Fatalf("Could not parse the response body %v", err)
	}

	if (&model.Product{}) == product {
		t.Fatal("Item is empty")
	}
}
