package global

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type PostStructs struct {
	Text     string `json:"text"`
	Category string `json:"cat"`
	Title    string `json:"title"`
}

type PostInfo struct {
	Iduser   string
	Category string
	PostText string
	Time     string
	Title    string
}

type GoodRequest struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Comment struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	Title    string `json:"title"`
	IdUser   string `json:"iduser"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Date     string `json:"date"`
	Pp       string `json:"pp"`
}

type Users struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Pp       string `json:"pp"`
	Cookie   string `json:"cookie"`
	Threads  string `json:"thread"`
	Limit    string `json:"limit"`
	Offset   string `json:"offset"`
}

type MessagePrivate struct {
	Type        string `json:"type"`
	Id          string `json:"id"`
	Username    string `json:"username"`
	Threads     string `json:"thread"`
	Message     string `json:"message"`
	Pp          string `json:"pp"`
	Date        string `json:"date"`
	LastMessage string `json:"lastmessage"`
}

type Chat struct {
	Username string `json:"username"`
	Cookie   string `json:"cookie"`
}

type CommentText struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type Notif struct {
	Id       int    `json:"id"`
	Content  string `json:"content"`
	IdUser   string `json:"iduser"`
	IdUser2  string `json:"iduser2"`
	Date     string `json:"date"`
	Username string `json:"username"`
}
