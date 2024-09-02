package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	global "real-time-rofu/back"
	"real-time-rofu/back/database"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type changeId struct {
	Type   string `json:"type"`
	IdRoom string `json:"idroom"`
	Name   string `json:"name"`
}

type Thread struct {
	Id      string `json:"id"`
	Iduser1 string `json:"iduser1"`
	Iduser2 string `json:"iduser2"`
}

type Notif struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Threads  string `json:"thread"`
	Message  string `json:"message"`
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {

	roomID := r.URL.Query().Get("room")
	if roomID == "" {
		http.Error(w, "Room ID is required", http.StatusBadRequest)
		return
	}

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m, roomID, database.LookingInDbThanksToCookie("name", r))
	m.addClient(roomID, client)

	go client.readMessages()
	go client.writeMessages()

}

func (m *Manager) addClient(roomID string, client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.rooms[roomID]; !ok {
		m.rooms[roomID] = make(ClientList)
	}
	m.rooms[roomID][client] = true
}

func (m *Manager) removeClient(roomID string, client *Client) {
	m.Lock()
	defer m.Unlock()
	if clients, ok := m.rooms[roomID]; ok {
		if _, ok := clients[client]; ok {
			client.connection.Close()
			delete(clients, client)
		}
		if len(clients) == 0 {
			delete(m.rooms, roomID)
		}
	}
}
func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c.roomID, c)
	}()

	for {
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket closed unexpectedly: %v", err)
			}
			break
		}

		var message Message
		err = json.Unmarshal(payload, &message)

		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		switch message.Type {
		case "text":
			c.broadcastMessage([]byte(message.Content))
			c.notifyOtherUser()
		case "click":
			split := strings.Split(message.Content, "/")
			c.changeIdOfTheOtherUser(split[0], split[1])
		case "istyping":
			c.isTypingFunction()
		default:
			log.Println("Unknown message type")
		}
	}
}

func (c *Client) isTypingFunction() {
	for _, roomClients := range c.manager.rooms {
		for wsclient := range roomClients {
			if wsclient.pseudo == database.LookingInDbThanksToWhatYouWant("name", "id", getTheOtherUserInDatabase(c.roomID, c.pseudo)) {
				wsclient.isTyping <- []byte("istyping")
			}
		}
	}
}

func (c *Client) changeIdOfTheOtherUser(idRoomUser string, nameOfTheOther string) {
	for _, roomClients := range c.manager.rooms {
		for wsclient := range roomClients {
			if wsclient.pseudo == nameOfTheOther {
				otherUser := database.LookingInDbThanksToWhatYouWant("name", "id", getTheOtherUserInDatabase(idRoomUser, wsclient.pseudo))

				message := changeId{
					Type:   "roomChange",
					IdRoom: idRoomUser,
					Name:   otherUser,
				}

				messageBytes, err := json.Marshal(message)
				if err != nil {
					fmt.Println(err)
					continue
				}
				wsclient.roomChange <- messageBytes
			}
		}
	}
}

func insertNotifInDb(user1 string, user2 string, typeOf string) {
	db := database.Opendb()
	defer db.Close()
	date := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("INSERT INTO `Notifs` (iduser, iduser2,content, date) VALUES (?,?, ?,?)", user1, user2, typeOf, date)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Client) broadcastMessage(payload []byte) {

	msg := global.MessagePrivate{
		Type:     "message",
		Message:  string(payload),
		Username: c.pseudo,
		Threads:  c.roomID,
		Pp:       database.LookingInDbThanksToWhatYouWant("pp", "name", c.pseudo),
	}
	database.InsertMessageInDb(string(payload), c.roomID, database.LookingInDbThanksToWhatYouWant("id", "name", c.pseudo))

	insertNotifInDb(database.LookingInDbThanksToWhatYouWant("id", "name", c.pseudo), getTheOtherUserInDatabase(c.roomID, c.pseudo), "sent a message")
	for wsclient := range c.manager.rooms[c.roomID] {

		messageBytes, err := json.Marshal(msg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		wsclient.egress <- messageBytes
	}
}

func (c *Client) notifyOtherUser() {
	otherUser := getTheOtherUserInDatabase(c.roomID, c.pseudo)
	for _, roomClients := range c.manager.rooms {
		for wsclient := range roomClients {
			if wsclient.pseudo == database.LookingInDbThanksToWhatYouWant("name", "id", otherUser) {
				if wsclient.roomID == "general" {
					wsclient.notif <- []byte("general")
				} else {
					wsclient.notif <- []byte(c.roomID)

				}
			}
		}
	}
}

type Online struct {
	Type  string   `json:"type"`
	Users []string `json:"users"`
}

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func (c *Client) writeMessages() {

	defer func() {
		c.manager.removeClient(c.roomID, c)
	}()

	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}
			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println(err)
			}
			var contentofMessage global.MessagePrivate
			err := json.Unmarshal(message, &contentofMessage)
			if err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

		case notifs, ok := <-c.notif:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}

			notif := Notif{
				Type:     "notif",
				Threads:  string(notifs),
				Username: c.pseudo,
			}

			msgMarshall, err := json.Marshal(notif)
			if err != nil {
				log.Print(err)
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, msgMarshall); err != nil {
				log.Println(err)
			}

		case roomChange, ok := <-c.roomChange:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, roomChange); err != nil {
				log.Println(err)
			}

		case isOnline, ok := <-c.isOnline:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}
			isOnlineStruct := Online{
				Type:  "online",
				Users: strings.Split(string(isOnline), ","),
			}

			msgMarshall, err := json.Marshal(isOnlineStruct)
			if err != nil {
				log.Print(err)
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, msgMarshall); err != nil {
				log.Println(err)
			}
		case _, ok := <-c.isTyping:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}

			isnotif := Notif{
				Type:    "istyping",
				Message: "is typing...",
			}

			msgMarshall, err := json.Marshal(isnotif)
			if err != nil {
				log.Print(err)
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, msgMarshall); err != nil {
				log.Println(err)
			}

		}

	}
}

func getTheOtherUserInDatabase(idThread string, name string) string {
	db := database.Opendb()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM `Threads` WHERE id = ?", idThread)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var thread Thread
	var empty string
	for rows.Next() {
		err = rows.Scan(&thread.Id, &thread.Iduser1, &thread.Iduser2, &empty)
		if err != nil {
			fmt.Println(err)
		}
	}
	if thread.Iduser1 == database.LookingInDbThanksToWhatYouWant("id", "name", name) {
		return thread.Iduser2
	}
	return thread.Iduser1

}
