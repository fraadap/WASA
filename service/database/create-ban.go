package database

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) NewBan(id int, userIDBanned int, timeStamp string) (int, error) {

	var banID = 0
	er1 := db.c.QueryRow("SELECT id FROM ban WHERE userID=? AND banned=?", id, userIDBanned).Scan(&banID)
	if errors.Is(er1, sql.ErrNoRows) {
		res, err := db.c.Exec("INSERT INTO ban (userID,banned,timestamp) VALUES (?,?,?)", id, userIDBanned, timeStamp)
		t, _ := res.LastInsertId()
		banID = int(t)

		return banID, err
	} else {
		er1 = sqlite3.ErrConstraintUnique
		return banID, er1
	}

}
