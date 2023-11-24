package database

import (
	"database/sql"
	"errors"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) SetUsername(id int, username string) error {

	var temp int
	er := db.c.QueryRow("SELECT id FROM user WHERE id=?", id).Scan(&temp)
	if errors.Is(er, sql.ErrNoRows) {
		return er
	}

	_, err := db.c.Exec("UPDATE user SET username = ? WHERE id = ?", username, id)

	return err
}
