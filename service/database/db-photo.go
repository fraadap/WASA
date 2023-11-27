package database

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) NewPhoto(id int, path string, timestamp string) (int, error) {

	var photoID = 0
	er1 := db.c.QueryRow("SELECT id FROM photo WHERE path=?", path).Scan(&photoID)
	if errors.Is(er1, sql.ErrNoRows) {

		res, err := db.c.Exec("INSERT INTO photo (userID,path,timestamp) VALUES (?,?,?)", id, path, timestamp)

		if err != nil {
			return photoID, err
		}

		t, _ := res.LastInsertId()
		photoID = int(t)

		return photoID, err
	} else {
		er1 = sqlite3.ErrConstraintUnique
		return photoID, er1
	}

}

func (db *appdbimpl) DeletePhoto(id int, photoId int) error {

	ris, err := db.c.Exec("DELETE FROM photo WHERE userID=? AND id=?", id, photoId)

	if err != nil {
		return err
	} else {
		rows, _ := ris.RowsAffected()
		if rows == 0 {
			err1 := sql.ErrNoRows
			return err1
		} else {
			db.c.Exec("DELETE FROM like WHERE photoID=", photoId)
			db.c.Exec("DELETE FROM comment WHERE photoID=", photoId)
		}
	}

	return nil

}

func (db *appdbimpl) UserIDByPhoto(photoID int) (int, error) {
	var ban int
	err := db.c.QueryRow("SELECT userID FROM photo WHERE userID=?", photoID).Scan(&ban)
	return ban, err
}
