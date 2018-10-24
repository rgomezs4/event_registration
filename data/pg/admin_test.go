package pg

import (
	"testing"

	"github.com/rgomezs4/event_registration/data/model"
)

var adminKey model.Key

func TestAdminInsert(t *testing.T) {
	a := Admin{}
	tx, err := db.Begin()
	checkError(t, err)

	admin := model.Admin{
		Username: "user1",
		Password: "123456",
		Name:     "Admin",
	}
	adminKey, err = a.Insert(tx, admin)
	checkError(t, err)

	if adminKey <= 0 {
		t.Fatalf("unable to insert user, got key with value %d", adminKey)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestAdminFind(t *testing.T) {
	a := Admin{}
	tx, err := db.Begin()
	checkError(t, err)

	admin, err := a.Find(tx, adminKey)
	checkError(t, err)

	if admin == nil {
		t.Fatal("unable to fetch admin")
	}

	if admin.ID != adminKey {
		t.Fatalf("expected admin id to be %d got %d", adminKey, admin.ID)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestAdminLogin(t *testing.T) {
	a := Admin{}
	tx, err := db.Begin()
	checkError(t, err)

	admin, err := a.Login(tx, "user1", "123456")
	checkError(t, err)

	if admin == nil {
		t.Fatal("unable to login")
	}

	if admin.Username != "user1" {
		t.Fatalf("expected admin username to be 'user1' got %s", admin.Username)
	}

	if admin.Password != "" {
		t.Fatalf("expected password to be empty got %s", admin.Password)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestAdminUpdate(t *testing.T) {
	a := Admin{}
	tx, err := db.Begin()
	checkError(t, err)

	admin := model.UpdateAdmin{
		Name: "Administrator",
	}

	ad, err := a.Update(tx, adminKey, admin)
	checkError(t, err)

	if ad.Name != admin.Name {
		t.Fatalf("unable to update user")
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

func TestAdminUpdatePassword(t *testing.T) {
	a := Admin{}
	tx, err := db.Begin()
	checkError(t, err)

	if err := a.UpdatePassword(tx, adminKey, "123456"); err != nil {
		t.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
