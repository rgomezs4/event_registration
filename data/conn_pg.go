package data

import (
	"github.com/rgomezs4/event_registration/data/model"
	"github.com/rgomezs4/event_registration/data/pg"
)

// Open - opens the connection to the pg database
func (db *DB) Open(driverName, dataSourceName string) error {
	conn, err := model.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	db.Admin = &pg.Admin{}
	db.Person = &pg.Person{}
	db.Item = &pg.Item{}
	db.Product = &pg.Product{}
	db.Order = &pg.Order{}

	db.DatabaseName = "events"
	db.Connection = conn
	return nil
}
