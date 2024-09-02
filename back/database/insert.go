package database

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func CreateChatInsertedInDB(w http.ResponseWriter, user string, user2 string) {
	db := Opendb()
	defer db.Close()
	timeNow := time.Now()
	time := timeNow.Format("2006-01-02 15:04:05")

	result, err := db.Exec(`INSERT INTO Threads (iduser, iduser2, date) VALUES (?, ?, ?)`, user, user2, time)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	threadID, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := ID{ID: threadID}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

type ID struct {
	ID int64 `json:"id"`
}

func InsertMessageInDb(message string, roomID string, user string) {
	db := Opendb()
	defer db.Close()
	timeNow := time.Now()
	time := timeNow.Format("2006-01-02 15:04:05")
	_, err := db.Exec(`INSERT INTO Messages (id_thread,content,iduser, date) VALUES (?, ?, ?,?)`, roomID, message, user, time)
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec(`UPDATE Threads SET date = ? WHERE id = ?`, time, roomID)
	if err != nil {
		log.Println(err)
	}

}

func InsertCommentInDb(iduser string, comment string, id string) {
	db := Opendb()
	defer db.Close()
	timeNow := time.Now()
	time := timeNow.Format("2006-01-02 15:04:05")
	_, err := db.Exec(`INSERT INTO Comments (id_thread, iduser, content, date) VALUES (?,?, ?, ?)`, id, iduser, comment, time)
	if err != nil {
		log.Println(err)
	}
}
