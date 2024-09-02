package database

import (
	"database/sql"
	"fmt"
	"net/http"
)

func LookingInDbThanksToCookie(valueToLookingFor string, r *http.Request) string {
	cookieValue := GetCookieValue(r)
	var storedValue string
	db := Opendb()
	defer db.Close()
	query := fmt.Sprintf(`SELECT %s FROM "User data" WHERE cookie = ?`, valueToLookingFor)
	err := db.QueryRow(query, cookieValue).Scan(&storedValue)
	if err != nil {
		if err == sql.ErrNoRows {
			storedValue = ""
		} else {
			fmt.Println(err)
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	return storedValue
}

func LookingInDbThanksToWhatYouWant(valueToLookingFor string, columYouWantToSearch string, valueYouHave string) string {
	var storedValue string
	db := Opendb()
	defer db.Close()
	query := fmt.Sprintf(`SELECT %s FROM "User data" WHERE %s = ?`, valueToLookingFor, columYouWantToSearch)
	err := db.QueryRow(query, valueYouHave).Scan(&storedValue)
	if err != nil {
		if err == sql.ErrNoRows {
			storedValue = ""
		} else {
			fmt.Println(err)
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	return storedValue
}

func FindThreadIdIndatabase(userid1 string, userid2 string) string {
	var threadId string
	db := Opendb()
	defer db.Close()
	query := `SELECT id FROM "Threads" WHERE (iduser = ? AND iduser2 = ?) OR (iduser = ? AND iduser2 = ?)`
	err := db.QueryRow(query, userid1, userid2, userid2, userid1).Scan(&threadId)
	if err != nil {
		if err == sql.ErrNoRows {
			threadId = "noThread"
		} else {
			threadId = "noThread"
		}
	}
	if err != nil {
		threadId = "noThread"
	}
	return threadId
}
