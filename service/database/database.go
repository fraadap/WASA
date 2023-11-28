/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/fraadap/WASA/service/structs"
	"github.com/sirupsen/logrus"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	Login(username string) (int, error)
	GetProfile(ID int) (structs.Profile, error)
	SetUsername(id int, username string) error
	NewFollow(id int, followedId int, timestamp string) (int, error)
	DeleteFollow(id int, followId int) error
	NewBan(id int, userIDBanned int, timeStamp string) (int, error)
	DeleteBan(id int, banId int) error
	IsBanned(userID int, bannedID int) (bool, error)
	NewPhoto(id int, path string, timestamp string) (int, error)
	DeletePhoto(id int, photoId int) error
	UserIDByPhoto(photoID int) (int, error)
	NewComment(userID int, photoID int, text string, TimeStamp string) (int, error)
	DeleteComment(commentID int, photoID int, userID int) error
	NewLike(userID int, photoID int, TimeStamp string) (int, error)
	DeleteLike(likeID int, photoID int, userID int) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	//db.Exec("DROP TABLE example_table")
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table'; `).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		err = createTables(db)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		logger.Println("Database's tables created")
	}

	db.Exec("PRAGMA foreign_keys = ON;") // abilitare le foreign keys

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
