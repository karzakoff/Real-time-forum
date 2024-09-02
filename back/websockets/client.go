package websockets

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn

	manager *Manager

	roomID string

	egress chan []byte

	notif chan []byte

	roomChange chan []byte

	pseudo string

	isOnline chan []byte

	isTyping chan []byte
}

func NewClient(conn *websocket.Conn, manager *Manager, roomID string, name string) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		roomID:     roomID,
		egress:     make(chan []byte),
		notif:      make(chan []byte),
		roomChange: make(chan []byte),
		isOnline:   make(chan []byte),
		isTyping:   make(chan []byte),

		pseudo: name,
	}
}

type Manager struct {
	rooms map[string]ClientList
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		rooms: make(map[string]ClientList),
	}
}
