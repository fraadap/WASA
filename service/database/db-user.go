package database

import (
	"database/sql"
	"errors"

	"github.com/fraadap/WASA/service/structs"
)

func (db *appdbimpl) Login(username string) (int, error) {
	var id int
	er1 := db.c.QueryRow("SELECT id FROM user WHERE username=?", username).Scan(&id)
	if errors.Is(er1, sql.ErrNoRows) {
		res, err := db.c.Exec("INSERT INTO user (username) VALUES (?)", username)
		if err != nil {
			return 0, err
		}
		t, _ := res.LastInsertId()
		id = int(t)
	} else if er1 != nil {
		return 0, er1
	}

	return id, nil
}

func (db *appdbimpl) SetUsername(id int, username string) error {

	var temp int
	er := db.c.QueryRow("SELECT id FROM user WHERE id=?", id).Scan(&temp)
	if errors.Is(er, sql.ErrNoRows) {
		return er
	}

	_, err := db.c.Exec("UPDATE user SET username = ? WHERE id = ?", username, id)

	return err
}

func (db *appdbimpl) GetProfile(ID int) (structs.Profile, error) {
	var profile structs.Profile

	// query per le info dell'utente
	queryUser := "SELECT id, username FROM user WHERE id = ?"
	err := db.c.QueryRow(queryUser, ID).Scan(&profile.User.Id, &profile.User.Username)
	if err != nil {
		return profile, err
	}

	// query per le foto dell'utente
	queryPhotos := "SELECT photo.id, photo.userID, photo.binary, photo.timestamp FROM photo WHERE photo.userID=? ORDER BY photo.timestamp DESC"
	photos, err := db.c.Query(queryPhotos, ID)
	if err != nil || photos.Err() != nil {
		return profile, err
	}

	// query per i followings dell'utente
	queryFollowings := "SELECT user.id, user.username FROM user, follow WHERE user.id = follow.followed AND follow.userID=?"
	followings, err := db.c.Query(queryFollowings, ID)
	if err != nil || followings.Err() != nil {
		return profile, err
	}

	// query per i followers dell'utente
	queryFollowers := "SELECT user.id, user.username FROM user, follow WHERE user.id=follow.userID AND follow.followed=?"
	followers, err := db.c.Query(queryFollowers, ID)
	if err != nil || followers.Err() != nil {
		return profile, err
	}

	// query per i ban dell'utente
	queryBans := "SELECT user.id, user.username FROM user, ban WHERE user.id=ban.banned AND ban.userID=?"
	bans, err := db.c.Query(queryBans, ID)
	if err != nil || bans.Err() != nil {
		return profile, err
	}

	// per ogni foto creo un tipo foto, per ogni foto prendo i commenti e i like
	for photos.Next() {
		var ph structs.PhotoInfo
		err := photos.Scan(&ph.Photo.PhotoID, &ph.Photo.UserID, &ph.Photo.Binary, &ph.Photo.TimeStamp)
		if err != nil {
			return profile, err
		}
		ph.Comments, err = db.GetComments(ph.Photo.PhotoID)
		if err != nil {
			return profile, err
		}

		ph.Likes, err = db.GetLikes(ph.Photo.PhotoID)
		if err != nil {
			return profile, err
		}

		ph.NComments = len(ph.Comments)
		ph.NLikes = len(ph.Likes)

		profile.Photos = append(profile.Photos, ph)
	}

	// per ogni foto creo un tipo foto, per ogni foto prendo i commenti e i like
	for followings.Next() {
		var u structs.User
		err1 := followings.Scan(&u.Id, &u.Username)
		if err1 != nil {
			return profile, err1
		} else {
			profile.Followings = append(profile.Followings, u)
		}
	}

	for followers.Next() {
		var u structs.User
		err1 := followers.Scan(&u.Id, &u.Username)
		if err1 != nil {
			return profile, err1
		} else {
			profile.Followers = append(profile.Followers, u)
		}
	}

	for bans.Next() {
		var u structs.User
		err1 := bans.Scan(&u.Id, &u.Username)
		if err1 != nil {
			return profile, err1
		} else {
			profile.Bans = append(profile.Bans, u)
		}
	}

	profile.NPhotos = len(profile.Photos)

	return profile, nil
}

func (db *appdbimpl) ExistsUser(userID int) (bool, error) {
	yes := false
	queryUser := "SELECT EXISTS(*) FROM user WHERE id = ?"
	err := db.c.QueryRow(queryUser, userID).Scan(&yes)
	if err != nil {
		return yes, err
	}
	return yes, err
}

func (db *appdbimpl) GetUsername(id int) (string, error) {
	var username string

	// query per le info dell'utente
	queryUser := "SELECT username FROM user WHERE id = ?"
	err := db.c.QueryRow(queryUser, id).Scan(&username)
	if err != nil {
		return username, err
	}
	return username, err
}
