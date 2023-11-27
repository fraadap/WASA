package database

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) NewLike(userID int, photoID int, TimeStamp string) (int, error) {

	var likeID = 0

	er1 := db.c.QueryRow("SELECT id FROM like WHERE photoID=? AND userID=?", photoID, userID).Scan(&photoID)
	if errors.Is(er1, sql.ErrNoRows) {

		res, err := db.c.Exec("INSERT INTO like (photoID,userID,timestamp) VALUES (?,?,?)", photoID, userID, TimeStamp)

		if err != nil {
			return likeID, err
		}

		t, _ := res.LastInsertId()
		likeID = int(t)

		return likeID, err
	} else {
		er1 = sqlite3.ErrConstraintUnique
		return likeID, er1
	}
}

func (db *appdbimpl) DeleteLike(likeID int, photoID int, userID int) error {

	ris, err := db.c.Exec("DELETE FROM like WHERE userID=? AND photoid=? AND id=?", userID, photoID, likeID)

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
