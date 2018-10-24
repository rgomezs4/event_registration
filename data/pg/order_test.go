package pg

import (
	"testing"

	"github.com/rgomezs4/event_registration/data/model"
)

func TestOrderHeaderInsert(t *testing.T) {
	o := Order{}
	tx, err := db.Begin()
	checkError(t, err)

	header := model.OrderHeader{
		PersonID:      1,
		PaymentMethod: 1,
		Total:         100,
		Comment:       "none",
		BtcAddress:    "address",
	}

	orderKey, err := o.InsertHeader(tx, adminKey, header)
	checkError(t, err)

	if orderKey <= 0 {
		t.Fatalf("unable to insert header, got key with value %d", orderKey)
	}

	if err := tx.Rollback(); err != nil {
		t.Fatal(err)
	}
}

func TestOrderDetailInsert(t *testing.T) {
	o := Order{}
	p := Product{}
	tx, err := db.Begin()
	checkError(t, err)

	product := model.Product{
		Name:    "Test Product",
		Barcode: "123456",
		Size:    "M",
		Stock:   1000,
		Price:   100.54,
	}

	productKey, err := p.Insert(tx, product)
	checkError(t, err)

	header := model.OrderHeader{
		PersonID:      1,
		PaymentMethod: 1,
		Total:         100,
		Comment:       "none",
		BtcAddress:    "address",
	}

	orderKey, err := o.InsertHeader(tx, adminKey, header)
	checkError(t, err)

	detail := model.OrderDetail{
		ProductID: productKey,
		Quantity:  10,
		Price:     10,
		Amount:    100,
	}

	key, err := o.InsertDetail(tx, orderKey, detail)
	checkError(t, err)

	if key <= 0 {
		t.Fatalf("unable to insert detail, got key with value %d", key)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestOrderGetByUserID(t *testing.T) {
	o := Order{}
	tx, err := db.Begin()
	checkError(t, err)

	orders, err := o.GetOrdersByID(tx, 1)
	checkError(t, err)

	if len(orders) <= 0 {
		t.Fatalf("should retrieve at least 1 order")
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestOrderGetDetail(t *testing.T) {
	o := Order{}
	tx, err := db.Begin()
	checkError(t, err)

	orders, err := o.FindDetail(tx, 2)
	checkError(t, err)

	if len(orders) <= 0 {
		t.Fatalf("should retrieve at least 1 order")
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestOrderGetTotals(t *testing.T) {
	o := Order{}
	tx, err := db.Begin()
	checkError(t, err)

	quantity, total, err := o.GetTotalsByUser(tx, 1)
	checkError(t, err)

	if quantity <= 0 {
		t.Fatalf("quantity must be > 0")
	}

	if total <= 0 {
		t.Fatalf("quantity must be > 0")
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
