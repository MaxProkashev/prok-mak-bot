package dbfunc

import (
	"database/sql"
	"log"
	"os"
)

// OpenDB - open connection with db by heroku env
func OpenDB() (db *sql.DB, err error) {
	return sql.Open("postgres", os.Getenv("DATABASE_URL"))
}

// CreateTable -
func CreateTable(db *sql.DB, name string) {
	if name == "bot_user" {
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + name + " (id INT PRIMARY KEY,name TEXT,surname TEXT,img TEXT,study TEXT,work TEXT,status TEXT,lastask INT,temp TEXT);")
		if err != nil {
			log.Fatalf("[X] Could not create %s table. Reason: %s", name, err.Error())
		} else {
			log.Printf("[OK] Create %s table", name)
		}
	} else if name == "asking" {
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + name + " (iduser INT,id INT PRIMARY KEY,idsolv INT,date TEXT,theme TEXT,info TEXT);")
		if err != nil {
			log.Fatalf("[X] Could not create %s table. Reason: %s", name, err.Error())
		} else {
			log.Printf("[OK] Create %s table", name)
		}
	} else {
		log.Printf("[ERR] Wrong %s table DB", name)
	}
}
