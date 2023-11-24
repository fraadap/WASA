package database

import (
	"database/sql"
)

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
