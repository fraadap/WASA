package database

import (
	"database/sql"
	"fmt"
)

func createTables(db *sql.DB) error {

	//	e := ""
	var err error
	var sqlStmt string

	//user table
	sqlStmt = `CREATE TABLE IF NOT EXISTS user (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, username VARCHAR(16));`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//follow table
	sqlStmt = `CREATE TABLE IF NOT EXISTS follow (id INTEGER NOT NULL PRIMARY KEY, userID INTEGER, followed INTEGER, timestamp DATETIME, FOREIGN KEY(userID) REFERENCES user (id), FOREIGN KEY(followed) REFERENCES user (id));`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//ban table
	sqlStmt = `CREATE TABLE IF NOT EXISTS ban (id INTEGER NOT NULL PRIMARY KEY, userID INTEGER, banned INTEGER, timestamp DATETIME, FOREIGN KEY(userID) REFERENCES user (id), FOREIGN KEY(banned) REFERENCES user (id));`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//photo table
	sqlStmt = `CREATE TABLE IF NOT EXISTS photo (id INTEGER NOT NULL PRIMARY KEY, userID INTEGER, path VARCHAR(40), timestamp DATETIME, FOREIGN KEY(userID) REFERENCES user (id));`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//comment table
	sqlStmt = `CREATE TABLE IF NOT EXISTS comment (id INTEGER NOT NULL PRIMARY KEY, userID INTEGER, photoID INTEGER, text VARCHAR(300), timestamp DATETIME, FOREIGN KEY(userID) REFERENCES user (id), FOREIGN KEY(photoID) REFERENCES photo (id));`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//like table
	sqlStmt = `CREATE TABLE IF NOT EXISTS like (id INTEGER NOT NULL PRIMARY KEY, userID INTEGER, photoID INTEGER, timestamp DATETIME, FOREIGN KEY(userID) REFERENCES user (id), FOREIGN KEY(photoID) REFERENCES photo (id));`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
