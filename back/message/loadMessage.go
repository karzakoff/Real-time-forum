package message

import (
	"encoding/json"
	"fmt"
	"net/http"
	global "real-time-rofu/back"
	"real-time-rofu/back/database"
	"strconv"
)

func LoadMessage(w http.ResponseWriter, r *http.Request, user global.Users) {
	var messages []global.MessagePrivate
	nbMessage := GetTheNumberOfMessageWithAThreadId(user.Threads)
	if nbMessage == 0 {

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(messages)
		return
	}

	Limit, _ := strconv.Atoi(user.Limit)

	db := database.Opendb()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM `Messages` WHERE id_thread = ? ORDER BY id ASC LIMIT ? OFFSET ?", user.Threads, nbMessage, nbMessage-Limit)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var message global.MessagePrivate
		err = rows.Scan(&message.Id, &message.Threads, &message.Message, &message.Username, &message.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		message.LastMessage = "false"
		message.Username = database.LookingInDbThanksToWhatYouWant("name", "id", message.Username)
		message.Pp = database.LookingInDbThanksToWhatYouWant("pp", "name", message.Username)
		messages = append(messages, message)
	}
	if len(messages) == nbMessage {
		messages[len(messages)-1].LastMessage = "true"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func GetTheNumberOfMessageWithAThreadId(thread_id string) int {
	db := database.Opendb()
	defer db.Close()
	rows, err := db.Query("SELECT COUNT(*) FROM `Messages` WHERE id_thread = ?", thread_id)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0
		}
	}
	return count
}
