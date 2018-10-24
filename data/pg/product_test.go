package pg

import (
	"testing"

	"github.com/rgomezs4/event_registration/data/model"
)

var productKey model.Key

func TestProductInsert(t *testing.T) {
	p := Product{}
	tx, err := db.Begin()
	checkError(t, err)

	product := model.Product{
		Name:    "Test Item",
		Barcode: "abc123",
		Size:    "M",
		Stock:   1000,
		Price:   100.54,
	}

	productKey, err := p.Insert(tx, product)
	checkError(t, err)

	if productKey <= 0 {
		t.Fatalf("unable to insert product, got key with value %d", adminKey)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestProductAll(t *testing.T) {
	p := Product{}
	tx, err := db.Begin()
	checkError(t, err)

	products, err := p.All(tx)
	checkError(t, err)

	if products == nil {
		t.Fatal("expected to fetch a product")
	}

	if len(products) <= 0 {
		t.Fatalf("should fetch at least 1 product got %d", len(products))
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestProductFind(t *testing.T) {
	p := Product{}
	tx, err := db.Begin()
	checkError(t, err)

	product, err := p.Find(tx, 1)
	checkError(t, err)

	if product == nil {
		t.Fatal("expected to fetch a product")
	}

	if product.ID != 1 {
		t.Fatalf("expected id to be 1 got %d", product.ID)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestProductFindByBarcode(t *testing.T) {
	p := Product{}
	tx, err := db.Begin()
	checkError(t, err)

	product, err := p.FindByBarcode(tx, "abc123")
	checkError(t, err)

	if product == nil {
		t.Fatal("expected to fetch a product")
	}

	if product.ID != 2 {
		t.Fatalf("expected id to be 2 got %d", product.ID)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestProductUpdate(t *testing.T) {
	p := Product{}
	tx, err := db.Begin()
	checkError(t, err)

	pr := model.Product{
		Name:    "Updated Item",
		Barcode: "updated_barcode",
		Size:    "L",
		Stock:   1000,
		Price:   10.1,
	}

	product, err := p.Update(tx, 1, pr)
	checkError(t, err)

	if product == nil {
		t.Fatal("expected to fetch a product after updating it")
	}

	if product.ID != 1 {
		t.Fatalf("expected id to be 1 got %d", product.ID)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
