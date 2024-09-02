package comment

import (
	"net/http"
	"real-time-rofu/back/database"
)

func RecupCommentAndPutInDb(r *http.Request, comment string, id string) {
	db := database.Opendb()
	defer db.Close()
	iduser := database.LookingInDbThanksToCookie("id", r)
	database.InsertCommentInDb(iduser, comment, id)
}


