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
