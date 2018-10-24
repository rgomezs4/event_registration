package pg

import (
	"database/sql"
	"github.com/rgomezs4/event_registration/data/model"
)

// Admin data service
type Admin struct {
}

// Insert - add a new user to the table
func (a *Admin) Insert(tx *sql.Tx, admin model.Admin) (model.Key, error) {
	var id model.Key
	sqlStatement := `INSERT INTO public.admins
	(username, "name", "password")
	VALUES($1, $2, crypt($3, gen_salt('bf')))	
	RETURNING id;`

	if err := tx.QueryRow(sqlStatement, admin.Username, admin.Name, admin.Password).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// Find fetches a user by its ID
func (a *Admin) Find(tx *sql.Tx, id model.Key) (*model.Admin, error) {
	admin := &model.Admin{}
	sqlStatement := `SELECT id, username, "name"
	FROM public.admins WHERE id = $1;`

	err := tx.QueryRow(sqlStatement, id).Scan(&admin.ID, &admin.Username, &admin.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return admin, nil
}

// Login checks for a user and password
func (a *Admin) Login(tx *sql.Tx, username, password string) (*model.Admin, error) {
	admin := &model.Admin{}
	sqlStatement := `SELECT id, username, "name"
	FROM public.admins WHERE username = $1 and password = crypt($2, password) LIMIT 1;`

	err := tx.QueryRow(sqlStatement, username, password).Scan(&admin.ID, &admin.Username, &admin.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return admin, nil
}

// Update updates the name of the admin
func (a *Admin) Update(tx *sql.Tx, id model.Key, admin model.UpdateAdmin) (ad *model.Admin, err error) {
	ad = &model.Admin{}
	sqlStatement := `UPDATE public.admins
	SET "name"=$1
	WHERE id=$2
	RETURNING id, username, "name";`

	err = tx.QueryRow(sqlStatement, admin.Name, id).Scan(&ad.ID, &ad.Username, &ad.Name)

	return
}

// UpdatePassword updates the password to a new encrypted password
func (a *Admin) UpdatePassword(tx *sql.Tx, id model.Key, password string) (err error) {
	sqlStatement := `UPDATE public.admins
	SET "password"=crypt($1, gen_salt('bf'))
	WHERE id=$2;`

	_, err = tx.Exec(sqlStatement, password, id)

	return
}
