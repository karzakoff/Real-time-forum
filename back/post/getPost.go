package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	global "real-time-rofu/back"
	"real-time-rofu/back/database"
	"time"
)

func GetPost(post global.PostStructs, r *http.Request, w http.ResponseWriter) {
	if insertPostIndb(recupInfo(r, post)) {
		sendValidation(w)
	}
}

func recupInfo(r *http.Request, postFromJs global.PostStructs) global.PostInfo {
	var postGoToDb global.PostInfo
	currentTime := time.Now()
	postGoToDb.Time = currentTime.Format("2006-01-02 15:04")
	postGoToDb.Iduser = database.LookingInDbThanksToCookie("id", r)
	postGoToDb.PostText = postFromJs.Text
	postGoToDb.Category = postFromJs.Category
	postGoToDb.Title = postFromJs.Title
	return postGoToDb
}

func insertPostIndb(post global.PostInfo) bool {
	db := database.Opendb()
	defer db.Close()
	_, err := db.Exec("INSERT INTO `Post` (`iduser`,'date', 'content', 'category', 'title') VALUES (?,?,?,?,?)", post.Iduser, post.Time, post.PostText, post.Category, post.Title)
	if err != nil {
		fmt.Println(" here ", err)
		return false
	} else {
		return true
	}
}

func sendValidation(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	goodresponse := global.GoodRequest{Message: "good"}
	json.NewEncoder(w).Encode(goodresponse)
}
