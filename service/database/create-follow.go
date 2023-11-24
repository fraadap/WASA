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
