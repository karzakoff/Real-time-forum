package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func DataBaseCreateIfNotExist() {
	_, err := os.Stat("./back/database/database.db")
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create("./back/database/database.db")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			createDatabase()

		}
	}
	db, err := sql.Open("sqlite3", "./back/database/database.db")
	if err != nil {
		fmt.Println(err, "error open database")
	}
	defer db.Close()

}
