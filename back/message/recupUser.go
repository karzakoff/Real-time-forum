package message

import (
	"encoding/json"
	"fmt"
	"net/http"
	global "real-time-rofu/back"
	"real-time-rofu/back/database"
	"sort"
	"unicode"
)

func RecupUser(w http.ResponseWriter, r *http.Request) {
	db := database.Opendb()
	defer db.Close()
	rows, err := db.Query("SELECT id FROM 'User data'")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	userWhoSent := database.LookingInDbThanksToCookie("name", r)
	var User []global.Users
	allthreads := recupEveryThreadsWhereThereIsTheUser(database.LookingInDbThanksToCookie("id", r))
	for rows.Next() {
		var user global.Users
		err = rows.Scan(&user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user.Username = database.LookingInDbThanksToWhatYouWant("name", "id", user.Id)
		user.Pp = database.LookingInDbThanksToWhatYouWant("pp", "id", user.Id)
		if user.Username != userWhoSent {
			user.Threads = database.FindThreadIdIndatabase(database.LookingInDbThanksToCookie("id", r), user.Id)
			User = append(User, user)

		}

	}

	UserWithThread, threadNumber := shortAllThreadWithId(allthreads, User)

	User = sortUsersByUsername(User)

	if threadNumber != 0 {
		NewUserSlice := make([]global.Users, 0)
		for _, k := range UserWithThread {
			for i, j := range User {
				if k.Threads == j.Threads {
					User = append(User[:i], User[i+1:]...)
				}
			}
		}

		NewUserSlice = append(NewUserSlice, UserWithThread...)
		NewUserSlice = append(NewUserSlice, User...)
		User = NewUserSlice
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User)
}
func recupEveryThreadsWhereThereIsTheUser(iduser string) []string {
	db := database.Opendb()
	defer db.Close()

	rows, err := db.Query(`
		SELECT id 
		FROM Threads 
		WHERE iduser = ? OR iduser2 = ? 
		ORDER BY date DESC`, iduser, iduser)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var threads []string
	for rows.Next() {
		var thread string
		err = rows.Scan(&thread)
		if err != nil {
			return nil
		}
		if lookIfthereareMessagesInThread(thread) {
			threads = append(threads, thread)
		}
	}
	return threads
}

func lookIfthereareMessagesInThread(threadid string) bool {
	db := database.Opendb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `Messages` WHERE id_thread = ?", threadid)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}
}

func shortAllThreadWithId(allthreadsid []string, User []global.Users) ([]global.Users, int) {
	var UserWithThread []global.Users
	var count int
	for _, k := range allthreadsid {
		for _, j := range User {
			if j.Threads == k {
				count++
				UserWithThread = append(UserWithThread, j)
			}
		}
	}
	return UserWithThread, count
}

func sortUsersByUsername(users []global.Users) []global.Users {
	sort.Slice(users, func(i, j int) bool {
		ui := users[i].Username
		uj := users[j].Username

		for k := 0; k < len(ui) && k < len(uj); k++ {
			ri := rune(ui[k])
			rj := rune(uj[k])

			if isSpecial(ri) && !isSpecial(rj) {
				return true
			}
			if !isSpecial(ri) && isSpecial(rj) {
				return false
			}
			if unicode.IsUpper(ri) && unicode.IsLower(rj) {
				return true
			}
			if unicode.IsLower(ri) && unicode.IsUpper(rj) {
				return false
			}
			if ri != rj {
				return ri < rj
			}
		}
		return len(ui) < len(uj)
	})
	return users
}
func isSpecial(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
}
