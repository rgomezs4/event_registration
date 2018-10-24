package pg

import (
	"database/sql"
	"github.com/rgomezs4/event_registration/data/model"
)

// Product data service
type Product struct {
}

// Insert add a new person to the table
func (p *Product) Insert(tx *sql.Tx, product model.Product) (model.Key, error) {
	var id model.Key
	sqlStatement := `INSERT INTO public.products
	("name", barcode, "size", stock, price)
	VALUES($1, $2, $3, $4, $5)
	returning id`

	if err := tx.QueryRow(sqlStatement, product.Name, product.Barcode, product.Size, product.Stock, product.Price).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// All fetches all the products
func (p *Product) All(tx *sql.Tx) ([]model.Product, error) {
	sqlStatement := `SELECT id, "name", barcode, "size", stock, price
	FROM public.products;`

	rows, err := tx.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		p := model.Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Barcode, &p.Size, &p.Stock, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// Find fetches a product by its ID
func (p *Product) Find(tx *sql.Tx, id model.Key) (*model.Product, error) {
	product := &model.Product{}
	sqlStatement := `SELECT id, "name", barcode, "size", stock, price
	FROM public.products	
	WHERE id = $1;`

	err := tx.QueryRow(sqlStatement, id).Scan(&product.ID, &product.Name, &product.Barcode, &product.Size, &product.Stock, &product.Price)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return product, nil
}

// FindByBarcode fetches a product by its barcode
func (p *Product) FindByBarcode(tx *sql.Tx, barcode string) (*model.Product, error) {
	product := &model.Product{}
	sqlStatement := `SELECT id, "name", barcode, "size", stock, price
	FROM public.products	
	WHERE barcode = $1
	LIMIT 1;`

	err := tx.QueryRow(sqlStatement, barcode).Scan(&product.ID, &product.Name, &product.Barcode, &product.Size, &product.Stock, &product.Price)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return product, nil
}

// Update updates a product by its id
func (p *Product) Update(tx *sql.Tx, id model.Key, product model.Product) (pr *model.Product, err error) {
	pr = &model.Product{}
	sqlStatement := `UPDATE public.products
	SET "name"=$2, barcode=$3, "size"=$4, stock=$5, price=$6
	WHERE id=$1
	returning id, "name", barcode, "size", stock, price;`

	err = tx.QueryRow(sqlStatement, id, product.Name, product.Barcode, product.Size, product.Stock, product.Price).
		Scan(&pr.ID, &pr.Name, &pr.Barcode, &pr.Size, &pr.Stock, &pr.Price)
	return
}
