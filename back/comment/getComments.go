package comment

import (
	"encoding/json"
	"fmt"
	"net/http"
	global "real-time-rofu/back"
	"real-time-rofu/back/database"
)

func GetCommFromDb(w http.ResponseWriter, r *http.Request, id string) {
	db := database.Opendb()
	defer db.Close()
	rows, err := db.Query("SELECT id, iduser, content, date FROM 'Comments' WHERE id_thread = ?", id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var comments []global.Comment
	for rows.Next() {
		var comment global.Comment
		err = rows.Scan(&comment.ID, &comment.IdUser, &comment.Content, &comment.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		comment.Username = database.LookingInDbThanksToWhatYouWant("name", "id", comment.IdUser)
		comment.Pp = database.LookingInDbThanksToWhatYouWant("pp", "id", comment.IdUser)
		comments = append(comments, comment)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
