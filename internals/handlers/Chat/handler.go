package chat

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/models"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ClientManager struct {
	Clients   map[string]map[*websocket.Conn]string // Grouped clients by groupId
	Broadcast chan models.Message
	Mutex     sync.Mutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:   make(map[string]map[*websocket.Conn]string), // Maps groupId -> clients
		Broadcast: make(chan models.Message),
	}
}

func (manager *ClientManager) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v\n", err)
		return
	}
	defer ws.Close()

	groupId := r.URL.Query().Get("groupId")
	if groupId == "" {
		log.Println("Group ID is missing")
		return
	}

	manager.Mutex.Lock()
	if _, exists := manager.Clients[groupId]; !exists {
		manager.Clients[groupId] = make(map[*websocket.Conn]string)
	}
	manager.Clients[groupId][ws] = ""
	manager.Mutex.Unlock()

	log.Printf("New client connected to group: %s", groupId)
	defer manager.HandleClientExit(groupId, ws)

	// Notify other clients about the new peer
	manager.broadcastNewPeer(groupId, ws)

	// Read incoming messages from the WebSocket
	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			break
		}
		log.Printf("Received message: %+v\n", msg)

		// Ensure the group ID is included in the message
		msg.Group = groupId
		manager.Broadcast <- msg
	}
}

func (manager *ClientManager) broadcastNewPeer(groupId string, ws *websocket.Conn) {
	// Broadcast NEW_PEER message to all clients in the group
	manager.Mutex.Lock()
	for client := range manager.Clients[groupId] {
		if client != ws {
			msg := models.Message{
				Type:   "NEW_PEER",
				Group:  groupId,
				PeerID: ws.RemoteAddr().String(), // You can adjust this to send a peer ID
			}
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error broadcasting NEW_PEER message to client: %v\n", err)
			}
		}
	}
	manager.Mutex.Unlock()
}

func (manager *ClientManager) HandleMessages() {
	for {
		msg := <-manager.Broadcast
		log.Printf("Broadcasting message to group: %s\n", msg.Group)

		manager.Mutex.Lock()
		for client := range manager.Clients[msg.Group] {
			if err := client.WriteJSON(msg); err != nil {
				log.Printf("Error sending message to client: %v\n", err)
				client.Close()
				delete(manager.Clients[msg.Group], client)
			} else {
				log.Printf("Message sent to client in group: %s\n", msg.Group)
			}
		}
		manager.Mutex.Unlock()
	}
}

func (manager *ClientManager) HandleClientExit(groupId string, ws *websocket.Conn) {
	manager.Mutex.Lock()
	if clients, exists := manager.Clients[groupId]; exists {
		delete(clients, ws)
		if len(clients) == 0 {
			delete(manager.Clients, groupId)
		}
	}
	manager.Mutex.Unlock()

	// Notify other clients that a peer has disconnected
	manager.broadcastRemovePeer(groupId, ws)
}

func (manager *ClientManager) broadcastRemovePeer(groupId string, ws *websocket.Conn) {
	// Broadcast REMOVE_PEER message to all clients in the group
	manager.Mutex.Lock()
	for client := range manager.Clients[groupId] {
		if client != ws {
			msg := models.Message{
				Type:   "REMOVE_PEER",
				Group:  groupId,
				PeerID: ws.RemoteAddr().String(), // Same as before, use appropriate peer ID
			}
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error broadcasting REMOVE_PEER message to client: %v\n", err)
			}
		}
	}
	manager.Mutex.Unlock()
}
