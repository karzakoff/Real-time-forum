package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//Xibqan-8mewni-vodteb

var databasePath = "./back/database/database.db"

func Opendb() *sql.DB {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func databaseCreateIfEmpty() bool {
	_, err := os.Stat(databasePath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(databasePath)
			if err != nil {
				log.Fatal(err)
				return false
			}
			createDatabase()
			fmt.Println("- database created -")
			defer file.Close()
		}
	}
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		fmt.Println(err, "error open database")
		return false
	}
	defer db.Close()
	return true
}

func createDatabase() {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		fmt.Println(err, "error open database")
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "User data" 
		(id INTEGER PRIMARY KEY,
		name TEXT UNIQUE,
		mail TEXT UNIQUE, 
		pp TEXT,
		passwordHashed TEXT,
		cookie TEXT UNIQUE)`)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "Post" 
		(id INTEGER PRIMARY KEY,
		iduser INTEGER,
		content TEXT, 
		category TEXT,
		date TEXT,
		title TEXT)`)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "Threads" 
		(id INTEGER PRIMARY KEY,
		iduser TEXT,
		iduser2 TEXT,
		date TEXT)`)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "Messages" 
		(id INTEGER PRIMARY KEY,
		id_thread INTEGER,
		content TEXT,
		iduser INT,
		date TEXT)`)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "Comments" 
		(id INTEGER PRIMARY KEY,
		id_thread INTEGER,
		content TEXT,
		iduser INT,
		date TEXT)`)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "Notifs" 
		(id INTEGER PRIMARY KEY,
		content TEXT,
		iduser INT,
		iduser2 INT,
		date TEXT)`)
	if err != nil {
		fmt.Println(err)
	}
	insertRandomUsers(db)
}

const (
	passwordHashed = "Xibqan-8mewni-vodteb"
	numRandomUsers = 20
	charset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func insertRandomUsers(db *sql.DB) {
	passwordHashede, err := HashPassword(passwordHashed)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < numRandomUsers; i++ {
		name := randomString(10)
		mail := fmt.Sprintf("%s@example.com", randomString(10))
		pp := "back/database/img/profile.png"
		_, err := db.Exec(`INSERT INTO "User data" (name, mail, pp, passwordHashed) VALUES (?, ?, ?, ?)`,
			name, mail, pp, passwordHashede)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("^^^^^^^^^^^^^")
		fmt.Println("inserted", name, mail, pp, passwordHashed)
		fmt.Println("============")
	}
	fmt.Printf("%d random users inserted.\n", numRandomUsers)
}

func randomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func LoadingDatabase() {
	if databaseCreateIfEmpty() {
		fmt.Println("-- database, good.")
	} else {
		fmt.Println("-- database, error.")
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
