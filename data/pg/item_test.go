package pg

import (
	"testing"
	"time"

	"github.com/rgomezs4/event_registration/data/model"
)

var itemKey model.Key

func TestItemInsertItem(t *testing.T) {
	i := Item{}
	tx, err := db.Begin()
	checkError(t, err)

	item := model.Item{
		Name: "testItem",
	}
	itemKey, err = i.InsertItem(tx, item)
	checkError(t, err)

	if itemKey <= 0 {
		t.Fatalf("unable to item, got key with value %d", itemKey)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestItemAllItems(t *testing.T) {
	i := Item{}
	tx, err := db.Begin()
	checkError(t, err)

	items, err := i.AllItems(tx)
	checkError(t, err)

	if len(items) <= 0 {
		t.Fatal("must retrieve some items")
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestItemInsertPersonItem(t *testing.T) {
	p := Person{}
	i := Item{}
	tx, err := db.Begin()
	checkError(t, err)

	person := model.Person{
		FirstName:      "Person",
		LastName:       "#1",
		Birthdate:      time.Now(),
		PassportNumber: "p#1",
		CountryOrigin:  "China",
		CountryBirth:   "Viet Nam",
		Language:       "English",
		Gender:         "Male",
		Transafer:      "100",
		MasterCouncil:  " ? ? ",
		Image:          "image10",
		Status:         model.StatusRegistered,
		Section:        "section 1",
		Position:       "position 3",
		Notes:          "arriving at 10pm",
		UpdatedBy:      adminKey,
	}

	personKey, err := p.Insert(tx, adminKey, person)
	checkError(t, err)

	item := model.PersonItem{
		PersonID: personKey,
		ItemID:   itemKey,
	}

	key, err := i.InsertPersonItem(tx, item, adminKey)
	checkError(t, err)

	if key <= 0 {
		t.Fatalf("unable to item, got key with value %d", key)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestPersonItems(t *testing.T) {
	i := Item{}
	tx, err := db.Begin()
	checkError(t, err)

	items, err := i.PersonItems(tx, 1)
	checkError(t, err)

	if len(items) <= 0 {
		t.Fatal("must retrieve some items")
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
