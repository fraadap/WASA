package database

import (
	"database/sql"
	"errors"
)

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

func (db *appdbimpl) SetUsername(id int, username string) error {

	var temp int
	er := db.c.QueryRow("SELECT id FROM user WHERE id=?", id).Scan(&temp)
	if errors.Is(er, sql.ErrNoRows) {
		return er
	}

	_, err := db.c.Exec("UPDATE user SET username = ? WHERE id = ?", username, id)

	return err
}
