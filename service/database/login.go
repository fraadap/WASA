package database

import (
	"database/sql"
	"errors"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Login(username string) (int, error) {
	var id int
	er1 := db.c.QueryRow("SELECT id FROM user WHERE username=?", username).Scan(&id)
	if errors.Is(er1, sql.ErrNoRows) {
		res, _ := db.c.Exec("INSERT INTO user (username) VALUES (?)", username)
		t, _ := res.LastInsertId()
		id = int(t)
	}

	return id, nil
}
