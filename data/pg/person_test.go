package pg

import (
	"testing"
	"time"

	"github.com/rgomezs4/event_registration/data/model"
)

var personKey model.Key

func TestPersonInsert(t *testing.T) {
	p := Person{}
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
		UpdatedBy:      1,
	}
	personKey, err = p.Insert(tx, 1, person)
	checkError(t, err)

	if personKey <= 0 {
		t.Fatalf("unable to insert user, got key with value %d", adminKey)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestPersonFind(t *testing.T) {
	p := Person{}
	tx, err := db.Begin()
	checkError(t, err)

	person, err := p.Find(tx, personKey)
	checkError(t, err)

	if person.ID != personKey {
		t.Fatalf("expected person id to be %d got %d", personKey, person.ID)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestPersonAll(t *testing.T) {
	p := Person{}
	tx, err := db.Begin()
	checkError(t, err)

	persons, err := p.All(tx)
	checkError(t, err)

	if len(persons) <= 0 {
		t.Fatalf("should retrieve at least 1 person got %d", len(persons))
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestPersonUpdate(t *testing.T) {
	p := Person{}
	tx, err := db.Begin()
	checkError(t, err)

	person := model.Person{
		FirstName:      "Asistee",
		LastName:       "One",
		Birthdate:      time.Now(),
		PassportNumber: "jp#1",
		CountryOrigin:  "Japan",
		CountryBirth:   "Korea",
		Language:       "Korean",
		Gender:         "Female",
		Transafer:      "?",
		MasterCouncil:  "?",
		Image:          "3.jpg",
		Status:         1,
		Section:        "147",
		Position:       "position 98",
		Notes:          "none",
		UpdatedBy:      1,
	}

	pe, err := p.Update(tx, personKey, 1, person)
	checkError(t, err)

	if pe.ID != personKey {
		t.Fatalf("expected key to be %d got %d", personKey, pe.ID)
	}

	if pe.FirstName != person.FirstName {
		t.Fatalf("expected updated firstname to be %s got %s", person.FirstName, pe.FirstName)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestPersonRegister(t *testing.T) {
	p := Person{}
	tx, err := db.Begin()
	checkError(t, err)

	err = p.Register(tx, 9, 3, "asistee_one.png")
	checkError(t, err)

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
