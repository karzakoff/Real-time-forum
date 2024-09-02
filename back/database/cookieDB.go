package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

func GetCookieValue(r *http.Request) string {
	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func PutCookieInDb(sessionToken string, name string) {
	db := Opendb()
	defer db.Close()
	query := "UPDATE `User data` SET `cookie` = ? WHERE UPPER(name) = ?"
	_, err := db.Exec(query, sessionToken, strings.ToUpper(name))
	if err != nil {
		fmt.Println(" cookie update error :  ", err)
		return
	}
}

func getCookie(r *http.Request) string {
	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		fmt.Println("error")
		return ""
	}
	return cookie.Value
}

func CheckCookieInDb(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(r)
	db := Opendb()
	defer db.Close()
	var query, storedName string
	query = "SELECT UPPER(name) FROM `User data` WHERE UPPER(cookie) = UPPER(?)"
	err := db.QueryRow(query, cookie).Scan(&storedName)
	if err != nil {
		if err == sql.ErrNoRows {
			LogoutCookie(w, r)
			return
		}
		LogoutCookie(w, r)
		return
	}
}

func LogoutCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "sessionToken",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
