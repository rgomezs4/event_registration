package pg

import (
	"database/sql"
	"github.com/rgomezs4/event_registration/data/model"
)

// Order data service
type Order struct {
}

// InsertHeader inserts the header of an order
func (o *Order) InsertHeader(tx *sql.Tx, userID model.Key, header model.OrderHeader) (model.Key, error) {
	var id model.Key
	sqlStatement := `INSERT INTO public.order_enc
	(person_id, " payment_method", total, "comment", btc_address, created_by)
	VALUES($1, $2, $3, $4, $5, $6)
	returning id;`

	if err := tx.QueryRow(sqlStatement, header.PersonID, header.PaymentMethod, header.Total, header.Comment, header.BtcAddress, userID).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// InsertDetail inserts the detail of the order
func (o *Order) InsertDetail(tx *sql.Tx, orderID model.Key, detail model.OrderDetail) (model.Key, error) {
	var id model.Key
	sqlStatement := `INSERT INTO public.order_det
	(oreder_enc_id, product_id, " quantity", price, amount)
	VALUES($1, $2, $3, $4, $5)
	returning id;`

	if err := tx.QueryRow(sqlStatement, orderID, detail.ProductID, detail.Quantity, detail.Price, detail.Amount).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetOrdersByID gets every order placed by an user
func (o *Order) GetOrdersByID(tx *sql.Tx, userID model.Key) ([]model.OrderHeader, error) {
	sqlStatement := `SELECT oe.id, oe.person_id, p.first_name, p.last_name, oe." payment_method", oe.total, oe."comment", oe.btc_address, oe.created_by
	FROM public.order_enc as oe
	inner join person as p on oe.person_id = p.id
	where created_by=$1;`

	rows, err := tx.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.OrderHeader
	for rows.Next() {
		o := model.OrderHeader{}
		err := rows.Scan(&o.ID, &o.PersonID, &o.FirstName, &o.LastName, &o.PaymentMethod, &o.Total, &o.Comment, &o.BtcAddress, &o.CreatedBy)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

// FindHeader finds a header by its id
func (o *Order) FindHeader(tx *sql.Tx, orderID model.Key) (*model.OrderHeader, error) {
	order := &model.OrderHeader{}
	sqlStatement := `SELECT oe.id, oe.person_id, p.first_name, p.last_name, oe." payment_method", oe.total, oe."comment", oe.btc_address, oe.created_by
	FROM public.order_enc as oe
	inner join person as p on oe.person_id = p.id
	where oe.id=$1;`

	err := tx.QueryRow(sqlStatement, orderID).Scan(&order.ID, &order.PersonID, &order.FirstName, &order.LastName, &order.PaymentMethod,
		&order.Total, &order.Comment, &order.BtcAddress, &order.CreatedBy)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return order, nil
}

// FindDetail fetches all of the details based on the order id
func (o *Order) FindDetail(tx *sql.Tx, orderID model.Key) ([]model.OrderDetail, error) {
	sqlStatement := `SELECT od.id, od.oreder_enc_id, od.product_id, p."name", p."size" ,od." quantity", od.price, od.amount
	FROM public.order_det as od
	inner join products as p on p.id = od.product_id
	where od.oreder_enc_id = $1;`

	rows, err := tx.Query(sqlStatement, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []model.OrderDetail
	for rows.Next() {
		o := model.OrderDetail{}
		err := rows.Scan(&o.ID, &o.OrderHeaderID, &o.ProductID, &o.Name, &o.Size, &o.Quantity, &o.Price, &o.Amount)
		if err != nil {
			return nil, err
		}
		details = append(details, o)
	}
	return details, nil
}

// GetTotalsByUser gets the totals by userid
func (o *Order) GetTotalsByUser(tx *sql.Tx, userID model.Key) (quantity int, total float64, err error) {
	sqlStatement := `select count(*) quantity, coalesce(sum(total),0) total
	from order_enc
	where created_by = $1`

	err = tx.QueryRow(sqlStatement, userID).Scan(&quantity, &total)

	return
}
