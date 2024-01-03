package database

import (
	"database/sql"

	"github.com/fraadap/WASA/service/structs"
)

func (db *appdbimpl) NewPhoto(id int, binary []byte, timestamp string) (int, error) {

	var photoID = 0

	res, err := db.c.Exec("INSERT INTO photo (userID,binary,timestamp) VALUES (?,?,?)", id, binary, timestamp)

	if err != nil {
		return photoID, err
	}

	t, _ := res.LastInsertId()
	photoID = int(t)

	return photoID, err

}

func (db *appdbimpl) DeletePhoto(id int, photoId int) error {

	_, err := db.c.Exec("DELETE FROM like WHERE photoID=?", photoId)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("DELETE FROM comment WHERE photoID=?", photoId)
	if err != nil {
		return err
	}

	ris, err := db.c.Exec("DELETE FROM photo WHERE id=?", photoId)

	if err != nil {
		return err
	} else {
		rows, _ := ris.RowsAffected()
		if rows == 0 {
			err1 := sql.ErrNoRows
			return err1
		} else {
			return nil
		}
	}

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
	queryPhotos := "SELECT photo.id, photo.userID, photo.binary, photo.timestamp FROM photo, follow WHERE photo.userID=follow.followed AND follow.userID=? ORDER BY photo.timestamp DESC"
	photos, err := db.c.Query(queryPhotos, id)
	if err != nil || photos.Err() != nil {
		return st, err
	}

	for photos.Next() {
		var ph structs.PhotoInfo
		err := photos.Scan(&ph.Photo.PhotoID, &ph.Photo.UserID, &ph.Photo.Binary, &ph.Photo.TimeStamp)
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
