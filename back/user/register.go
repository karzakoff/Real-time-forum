package userFunctions

import (
	"database/sql"
	"fmt"
	global "real-time-rofu/back"
	"real-time-rofu/back/database"

	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func RegisterTesting(User *global.User, w http.ResponseWriter, r *http.Request) {
	if !checkingInDatabase("name", User.Username) {
		sendErrorResponse(w, "name already taken")
		return
	}
	if !checkingInDatabase("mail", User.Mail) {
		sendErrorResponse(w, "mail already taken")
		return
	}
	if !CheckEachRegex(User) {
		sendErrorResponse(w, "one of the input didn't respect rules")
		return
	}
	createUserInDatabase(User)
	loginUser(w, User.Username)

}

func createUserInDatabase(User *global.User) {
	var err error
	User.Password, err = HashPassword(User.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	db := database.Opendb()
	_, err = db.Exec("INSERT INTO `User data` (`name`, 'mail', 'pp', passwordHashed) VALUES (?,?,?,?)", User.Username, User.Mail, ".././back/database/img/blank-profile-picture-973460_1280.png", User.Password)
	if err != nil {
		fmt.Println(" here ", err)
		return
	}

}

func checkingInDatabase(column string, value string) bool {
	db := database.Opendb()
	defer db.Close()
	var query, storedName string
	query = fmt.Sprintf("SELECT UPPER(%s) FROM `User data` WHERE UPPER(%s) = UPPER(?)", column, column)
	err := db.QueryRow(query, value).Scan(&storedName)
	if err != nil {
		if err == sql.ErrNoRows {
			return true
		}
		fmt.Println(err)
		return true
	}

	return false
}

func CheckEachRegex(User *global.User) bool {
	if !validateUsername(User.Username) {
		return false
	}

	if !validateMail(User.Mail) {
		return false
	}

	if !validatePassword(User.Password) {
		return false
	}

	return true
}

// All functions for register.
func validateUsername(username string) bool {
	match, _ := regexp.MatchString("^[0-9a-zA-Z]{4,16}$", username)
	return match
}

func validateMail(mail string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", mail)
	return match
}

func validatePassword(password string) bool {
	match, _ := regexp.MatchString("^[A-Za-z0-9].{8,}$", password)
	return match
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
