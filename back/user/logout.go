package userFunctions

import (
	"net/http"
)

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
