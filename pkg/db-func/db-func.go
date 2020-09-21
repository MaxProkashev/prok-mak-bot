package dbfunc

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq" // pq driver for database/sql
)

// OpenDB - open connection with db by heroku env
func OpenDB() (db *sql.DB, err error) {
	return sql.Open("postgres", os.Getenv("DATABASE_URL"))
}

// CreateTable - bot_user (id INT PRIMARY KEY)
func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS bot_user (id INT PRIMARY KEY);")
	if err != nil {
		log.Fatalf("[X] Could not create bot_user table. Reason: %s", err.Error())
	} else {
		log.Println("[OK] Create bot_user table")
	}
}

// NewID - adding a user to the table if there is none
func NewID(db *sql.DB, userID int) {
	_, err := db.Exec("INSERT INTO bot_user (id) VALUES (" + strconv.Itoa(userID) + ");")
	if err != nil {
		log.Fatalf("[X] Could not insert user in bot_user. Reason: %s", err.Error())
	} else {
		log.Printf("[OK] New user %d", userID)
	}
}
