package server

import (
	"encoding/json"
	global "real-time-rofu/back"
	"real-time-rofu/back/comment"
	"real-time-rofu/back/database"
	"real-time-rofu/back/message"
	"real-time-rofu/back/post"
	userFunctions "real-time-rofu/back/user"
	"strconv"

	"net/http"
)

func LaunchWebPAge(myhttp *http.ServeMux) {

	myhttp.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	myhttp.Handle("/back/database/img/", http.StripPrefix("/back/database/img/", http.FileServer(http.Dir("./back/database/img"))))
	fs := http.FileServer(http.Dir("./static"))
	myhttp.Handle("/", fs)

	myhttp.HandleFunc("/register", registerHandler)
	myhttp.HandleFunc("/login", loginHandler)
	myhttp.HandleFunc("/logout", logoutHandler)
	myhttp.HandleFunc("/createPost", createPost)
	myhttp.HandleFunc("/recupPost", post.GetPostFromDb)
	myhttp.HandleFunc("/sentComment", sentComment)
	myhttp.HandleFunc("/recupComments", recupComment)
	myhttp.HandleFunc("/recupUser", message.RecupUser)
	myhttp.HandleFunc("/recupNotif", recupNotif)
	myhttp.HandleFunc("/loadMessage", loadMessageFromJs)
	myhttp.HandleFunc("/sendMessagePrivate", sendMessagePrivate)
	myhttp.HandleFunc("/createChat", createChat)
	myhttp.HandleFunc("/checkCookie", database.CheckCookieInDb)

}

func recupNotif(w http.ResponseWriter, r *http.Request) {
	message.RecupNotif(w, r)
}

func recupComment(w http.ResponseWriter, r *http.Request) {
	var comments global.CommentText
	if err := json.NewDecoder(r.Body).Decode(&comments); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment.GetCommFromDb(w, r, strconv.Itoa(comments.Id))
}

func sentComment(w http.ResponseWriter, r *http.Request) {
	var comments global.CommentText
	if err := json.NewDecoder(r.Body).Decode(&comments); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment.RecupCommentAndPutInDb(r, comments.Text, strconv.Itoa(comments.Id))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("good")
}

func createChat(w http.ResponseWriter, r *http.Request) {
	var chat global.Chat

	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.CreateChatInsertedInDB(w, database.LookingInDbThanksToCookie("id", r), database.LookingInDbThanksToWhatYouWant("id", "name", chat.Username))
}

func sendMessagePrivate(w http.ResponseWriter, r *http.Request) {
	var msg global.MessagePrivate
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func loadMessageFromJs(w http.ResponseWriter, r *http.Request) {
	var user global.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message.LoadMessage(w, r, user)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var PostStruct global.PostStructs
	if err := json.NewDecoder(r.Body).Decode(&PostStruct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.GetPost(PostStruct, r, w)

}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user global.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userFunctions.RegisterTesting(&user, w, r)
}

type CookieStruct struct {
	Cookie string `json:"cookie"`
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	var cookie CookieStruct
	if err := json.NewDecoder(r.Body).Decode(&cookie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userFunctions.LogoutCookie(w, r)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user global.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userFunctions.LoginUserTesting(w, &user, r)
}
