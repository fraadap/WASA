package database

import (
	"database/sql"
	"errors"

	"github.com/fraadap/WASA/service/structs"
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
			_, err = db.c.Exec("DELETE FROM like WHERE photoID=", photoId)
			if err != nil {
				return err
			}
			_, err = db.c.Exec("DELETE FROM comment WHERE photoID=", photoId)
			if err != nil {
				return err
			}
		}
	}

	return nil

}

func (db *appdbimpl) UserIDByPhoto(photoID int) (int, error) {
	var id int
	err := db.c.QueryRow("SELECT userID FROM photo WHERE id=?", photoID).Scan(&id)
	return id, err
}

func (db *appdbimpl) GetMyStream(id int) (structs.Stream, error) {
	var st structs.Stream

	// query per le info dell'utente
	queryUser := "SELECT id, username FROM user WHERE id = ?"
	err := db.c.QueryRow(queryUser, id).Scan(&st.User.Id, &st.User.Username)
	if err != nil {
		return st, err
	}

	// query per le foto dell'utente
	queryPhotos := "SELECT photo.id, photo.userID, photo.path, photo.timestamp FROM photo, follow WHERE photo.userID=follow.followed AND follow.userID=? ORDER BY photo.timestamp DESC"
	photos, err := db.c.Query(queryPhotos, id)
	if err != nil {
		return st, err
	}

	for photos.Next() != false {
		var ph structs.PhotoInfo
		err := photos.Scan(&ph.Photo.PhotoID, &ph.Photo.UserID, &ph.Photo.Path, &ph.Photo.TimeStamp)
		if err != nil {
			return st, err
		}
		ph.Comments, err = db.GetComments(ph.Photo.PhotoID)
		if err != nil {
			return st, err
		}

		ph.Likes, err = db.GetLikes(ph.Photo.PhotoID)
		if err != nil {
			return st, err
		}

		ph.NComments = len(ph.Comments)
		ph.NLikes = len(ph.Likes)

		st.Photos = append(st.Photos, ph)
	}
	return st, nil
}
