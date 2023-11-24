package database

import (
	"database/sql"
)

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
