package websockets

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func SetupWebSockets(myhttp *http.ServeMux) {
	m := NewManager()
	myhttp.HandleFunc("/ws", m.serveWS)
	go loopManager(m)
}

func loopManager(m *Manager) {
	for {
		time.Sleep(500 * time.Millisecond)
		var allUserOnline []string
		for _, k := range m.rooms {
			for wsclient := range k {
				allUserOnline = append(allUserOnline, wsclient.pseudo)
			}
		}
		for _, k := range m.rooms {
			for wsclient := range k {
				wsclient.isOnline <- []byte(strings.Join(allUserOnline, ","))
			}
		}
	}
}

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)
