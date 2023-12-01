package database

import (
	"database/sql"

	"github.com/fraadap/WASA/service/structs"
)

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

func (db *appdbimpl) GetComments(photoID int) ([]structs.Comment, error) {
	var comments []structs.Comment

	queryUser := "SELECT * FROM comment WHERE photoID = ?"
	comms, err := db.c.Query(queryUser, photoID)
	if err != nil || comms.Err() != nil {
		return comments, err
	}
	for comms.Next() {
		var c structs.Comment
		err := comms.Scan(&c.CommentID, &c.UserID, &c.PhotoID, &c.Text, &c.TimeStamp)
		if err != nil {
			return comments, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func (db *appdbimpl) GetOwnerFromCommentID(commentID int) (int, error) {
	var owner int
	queryUser := "SELECT userID FROM comment WHERE id = ?"
	err := db.c.QueryRow(queryUser, commentID).Scan(&owner)
	return owner, err
}
