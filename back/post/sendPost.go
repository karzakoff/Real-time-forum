package post

import (
	"encoding/json"
	"net/http"
	global "real-time-rofu/back"
	"real-time-rofu/back/database"
)

func GetPostFromDb(w http.ResponseWriter, r *http.Request) {
	db := database.Opendb()
	defer db.Close()
	rows, err := db.Query("SELECT id, iduser, content, category, date, title FROM 'Post' ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var comments []global.Comment
	for rows.Next() {
		var comment global.Comment
		err = rows.Scan(&comment.ID, &comment.IdUser, &comment.Content, &comment.Category, &comment.Date, &comment.Title)
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
