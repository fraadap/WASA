package database

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) NewFollow(id int, followedId int, timestamp string) (int, error) {

	var followID = 0
	er1 := db.c.QueryRow("SELECT id FROM follow WHERE userID=? AND followed=?", id, followedId).Scan(&followID)
	if errors.Is(er1, sql.ErrNoRows) {
		res, err := db.c.Exec("INSERT INTO follow (userID,followed,timestamp) VALUES (?,?,?)", id, followedId, timestamp)
		t, _ := res.LastInsertId()
		followID = int(t)

		return followID, err
	} else {
		er1 = sqlite3.ErrConstraintUnique
		return followID, er1
	}

}

func (db *appdbimpl) DeleteFollow(id int, followId int) error {

	ris, err := db.c.Exec("DELETE FROM follow WHERE userID=? AND id=?", id, followId)

	if err != nil {
		return err
	} else {
		rows, _ := ris.RowsAffected()
		if rows == 0 {
			err1 := sql.ErrNoRows
			return err1
		}
	}

	return nil

}
