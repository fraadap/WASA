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

func (db *appdbimpl) DeleteBan(id int, banID int) error {

	ris, err := db.c.Exec("DELETE FROM ban WHERE userID=? AND id=?", id, banID)

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

func (db *appdbimpl) IsBanned(userID int, bannedID int) (bool, error) {
	var ban bool
	err := db.c.QueryRow("SELECT COUNT(*)=1 FROM ban WHERE userID=? AND banned=?", userID, bannedID).Scan(&ban)
	return ban, err
}

func (db *appdbimpl) GetBanID(id int, banned int) (int, error) {

	var banID = 0
	er1 := db.c.QueryRow("SELECT id FROM ban WHERE userID=? AND banned=?", id, banned).Scan(&banID)
	if er1 != nil {
		return 0, er1
	} else {
		return banID, nil
	}
}
