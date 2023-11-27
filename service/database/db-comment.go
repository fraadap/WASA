package database

import "database/sql"

func (db *appdbimpl) NewComment(userID int, photoID int, text string, TimeStamp string) (int, error) {
	var commentID = 0

	res, err := db.c.Exec("INSERT INTO comment (photoID,userID,text,timestamp) VALUES (?,?,?,?)", photoID, userID, text, TimeStamp)

	if err != nil {
		return commentID, err
	}

	t, _ := res.LastInsertId()
	commentID = int(t)

	return commentID, err

}

func (db *appdbimpl) DeleteComment(commentID int, photoID int, userID int) error {

	ris, err := db.c.Exec("DELETE FROM comment WHERE userID=? AND photoid=? AND id=?", userID, photoID, commentID)

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
