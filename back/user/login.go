package userFunctions

import (
	"database/sql"

	global "real-time-rofu/back"
	"real-time-rofu/back/database"

	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginUserTesting(w http.ResponseWriter, User *global.User, r *http.Request) {
	checkingPasswordAndName(w, *User)
}

func loginUser(w http.ResponseWriter, name string) {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(5000000 * time.Second)
	http.SetCookie(w, &http.Cookie{
		Name:    "sessionToken",
		Value:   sessionToken,
		Expires: expiresAt,
	})
	database.PutCookieInDb(sessionToken, name)
}

func checkingPasswordAndName(w http.ResponseWriter, User global.User) {
	var err error
	var passwordHash, query string

	db := database.Opendb()
	defer db.Close()

	if strings.Contains(User.Username, "@") {
		query = "SELECT passwordHashed FROM `User data` WHERE UPPER(mail) = ?"
	} else {
		query = "SELECT passwordHashed FROM `User data` WHERE UPPER(name) = ?"
	}

	row := db.QueryRow(query, strings.ToUpper(User.Username))
	err = row.Scan(&passwordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorResponse(w, "nom ou email pas trouvé")
			return
		} else {
			sendErrorResponse(w, "error")
			return
		}
	}
	if CheckPasswordHash(User.Password, passwordHash) {
		loginUser(w, User.Username)
		return
	} else {
		sendErrorResponse(w, "mot de passe pas bon")
		return
	}

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
