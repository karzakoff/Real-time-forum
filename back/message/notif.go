package message

import (
	"encoding/json"
	"net/http"
	global "real-time-rofu/back"
	"real-time-rofu/back/database"
)

func RecupNotif(w http.ResponseWriter, r *http.Request) {
	iduser := database.LookingInDbThanksToCookie("id", r)
	db := database.Opendb()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM `Notifs` WHERE iduser2 = ?", iduser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var notifs []global.Notif
	for rows.Next() {
		var notif global.Notif
		err = rows.Scan(&notif.Id, &notif.Content, &notif.IdUser, &notif.IdUser2, &notif.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		notif.Username = database.LookingInDbThanksToWhatYouWant("name", "id", notif.IdUser)
		notifs = append(notifs, notif)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifs)

}
