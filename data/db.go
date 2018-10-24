package data

import (
	"github.com/rgomezs4/event_registration/data/model"
	"github.com/rgomezs4/event_registration/data/pg"
)

// DB struct for our API
type DB struct {
	DatabaseName string
	Connection   *model.Connection

	Admin   *pg.Admin
	Person  *pg.Person
	Item    *pg.Item
	Product *pg.Product
	Order   *pg.Order
}
