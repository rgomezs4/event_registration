package pg

import (
	"database/sql"
	"github.com/rgomezs4/event_registration/data/model"
)

// Person data service
type Person struct {
}

// Insert add a new person to the table
func (p *Person) Insert(tx *sql.Tx, userID model.Key, person model.Person) (model.Key, error) {
	var id model.Key
	sqlStatement := `INSERT INTO public.person
	(first_name, last_name, birthdate, passport_number, country_origin, country_birth, "language", gender, transafer, mastercouncil, image, status, "section", "position", notes, updated_by)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	RETURNING id;`

	if err := tx.QueryRow(sqlStatement, person.FirstName, person.LastName, person.Birthdate, person.PassportNumber,
		person.CountryOrigin, person.CountryBirth, person.Language, person.Gender, person.Transafer, person.MasterCouncil,
		"", model.StatusPending, person.Section, person.Position, person.Notes, userID).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// Find fetches a person by its id
func (p *Person) Find(tx *sql.Tx, id model.Key) (*model.Person, error) {
	person := &model.Person{}
	sqlStatement := `SELECT id, first_name, last_name, birthdate, passport_number, country_origin, country_birth, "language", gender, transafer, mastercouncil, image, status, "section", "position", notes, updated_by
	FROM public.person
	WHERE id = $1;`

	err := tx.QueryRow(sqlStatement, id).Scan(&person.ID, &person.FirstName, &person.LastName, &person.Birthdate, &person.PassportNumber,
		&person.CountryOrigin, &person.CountryBirth, &person.Language, &person.Gender, &person.Transafer, &person.MasterCouncil,
		&person.Image, &person.Status, &person.Section, &person.Position, &person.Notes, &person.UpdatedBy)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return person, nil
}

// All fetches all the registries on the table
func (p *Person) All(tx *sql.Tx) ([]model.Person, error) {
	sqlStatement := `SELECT id, first_name, last_name, birthdate, passport_number, country_origin, country_birth, "language", gender, transafer, mastercouncil, image, status, "section", "position", notes, updated_by
	FROM public.person`

	rows, err := tx.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []model.Person
	for rows.Next() {
		p := model.Person{}
		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Birthdate, &p.PassportNumber,
			&p.CountryOrigin, &p.CountryBirth, &p.Language, &p.Gender, &p.Transafer, &p.MasterCouncil,
			&p.Image, &p.Status, &p.Section, &p.Position, &p.Notes, &p.UpdatedBy)
		if err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, nil
}

// Update updates the data of a person given its id
func (p *Person) Update(tx *sql.Tx, id, userID model.Key, person model.Person) (pe *model.Person, err error) {
	pe = &model.Person{}
	sqlStatement := `UPDATE public.person
	SET first_name=$2, last_name=$3, birthdate=$4, passport_number=$5, country_origin=$6, country_birth=$7, 
		"language"=$8, gender=$9, transafer=$10, mastercouncil=$11, "section"=$12, "position"=$13, notes=$14, updated_by=$15
	WHERE id=$1
	returning id, first_name, last_name, birthdate, passport_number, country_origin, country_birth, "language", gender, transafer, mastercouncil, image, status, "section", "position", notes, updated_by;`

	err = tx.QueryRow(sqlStatement, id, person.FirstName, person.LastName, person.Birthdate, person.PassportNumber, person.CountryOrigin,
		person.CountryBirth, person.Language, person.Gender, person.Transafer, person.MasterCouncil,
		person.Section, person.Position, person.Notes, userID).
		Scan(&pe.ID, &pe.FirstName, &pe.LastName, &pe.Birthdate, &pe.PassportNumber, &pe.CountryOrigin,
			&pe.CountryBirth, &pe.Language, &pe.Gender, &pe.Transafer, &pe.MasterCouncil, &pe.Image, &pe.Status,
			&pe.Section, &pe.Position, &pe.Notes, &pe.UpdatedBy)
	return
}

// Register change status and set image name
func (p *Person) Register(tx *sql.Tx, id, userID model.Key, image string) (err error) {
	sqlStatement := `UPDATE public.person
	SET image=$1, status=1, updated_by=$2
	WHERE id=$3;`
	_, err = tx.Exec(sqlStatement, image, userID, id)
	return
}
