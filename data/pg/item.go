package pg

import (
	"database/sql"
	"github.com/rgomezs4/event_registration/data/model"
)

// Item data service
type Item struct {
}

// InsertItem add new item to the table
func (i *Item) InsertItem(tx *sql.Tx, item model.Item) (model.Key, error) {
	var id model.Key
	sqlStatement := `INSERT INTO public.item
	("name")
	VALUES($1)
	RETURNING id;`

	if err := tx.QueryRow(sqlStatement, item.Name).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// AllItems get all items
func (i *Item) AllItems(tx *sql.Tx) ([]model.Item, error) {
	sqlStatement := `SELECT id, "name"
	FROM public.item;`

	rows, err := tx.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		i := model.Item{}
		err := rows.Scan(&i.ID, &i.Name)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

// InsertPersonItem inserts an item associated to a person
func (i *Item) InsertPersonItem(tx *sql.Tx, item model.PersonItem, userID model.Key) (model.Key, error) {
	var id model.Key
	sqlStatement := `INSERT INTO public.person_item
	(person_id, item_id, created_by)
	VALUES($1, $2, $3)		
	RETURNING id;`

	if err := tx.QueryRow(sqlStatement, item.PersonID, item.ItemID, userID).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// PersonItems get person associated items by personID
func (i *Item) PersonItems(tx *sql.Tx, personID model.Key) ([]model.PersonItem, error) {
	sqlStatement := `SELECT pi.id, pi.person_id, pi.item_id, i.name, pi.created_by
	FROM public.person_item as pi
	inner join public.item as i on i.id = pi.item_id
	where person_id = $1;`

	rows, err := tx.Query(sqlStatement, personID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.PersonItem
	for rows.Next() {
		i := model.PersonItem{}
		err := rows.Scan(&i.ID, &i.PersonID, &i.ItemID, &i.ItemName, &i.CreatedBy)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}
