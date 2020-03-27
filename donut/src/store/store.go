package store

import (
	"database/sql"

	"app/gen/donut"
)

// Store is a wrapper around a *sql.DB type for implementing interfaces using the DB.
type Store struct {
	DB *sql.DB
}

// GetDonutList gets the list of donuts from the database.
func (store Store) GetDonutList() (list donut.DonutList, err error) {
	rows, err := store.DB.Query("SELECT * FROM Donut;")
	defer rows.Close()
	if err != nil {
		return
	}
	
	for rows.Next() {
		var shape donut.Shape
		if err = rows.Scan(&shape); err != nil {
			return
		}
		donut := donut.Donut{Shape: shape}
		list.Donuts = append(list.Donuts, &donut)
	}
	
	if err = rows.Err(); err != nil {
		return
	}

	return
}

// AddDonut adds a donut to the database.
func (store Store) AddDonut(donut donut.Donut) error {
	rows, err := store.DB.Query("INSERT INTO Donut VALUES (1)")
	defer rows.Close()
	return err
}
