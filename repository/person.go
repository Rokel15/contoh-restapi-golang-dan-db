package repository

import (
	"database/sql"
	"mini-project-practice-build-rest-api/entities"
)

func GetAllPerson(db *sql.DB) (result []entities.Person, err error) {
	sql := "SELECT * from person"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var person entities.Person
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		if err != nil {
			return
		}
		result = append(result, person)
	}
	return
}

func InsertPerson(db *sql.DB, person entities.Person) (err error) {
	sql := "INSERT INTO person(id, first_name, last_name) values($1, $2, $3)"

	errs := db.QueryRow(sql, person.ID, person.FirstName, person.LastName)

	return errs.Err()
}

func UpdatePerson(db *sql.DB, person entities.Person) (err error) {
	sql := "UPDATE person SET first_name = $1, last_name = $2 WHERE id = $3"

	errs := db.QueryRow(sql, person.FirstName, person.LastName, person.ID)

	return errs.Err()
}

func DeletePerson(db *sql.DB, person entities.Person) (err error) {
	sql := "DELETE FROM PERSON WHERE id = $1"

	errs := db.QueryRow(sql, person.ID)

	return errs.Err()
}
