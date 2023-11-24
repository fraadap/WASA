package database

import (
	"database/sql"
)

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
