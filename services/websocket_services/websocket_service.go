package services

import (
	bc "bebuhcon/controller"
	"bebuhcon/utils"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	Bc *bc.BridgeController
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// func start listen http service
func (wss *WebSocketService) Start() {
	log.Printf("WebSocketService::Start executed")
	http.HandleFunc("/ws", wss.wsEndpoint)
	errServeWs := http.ListenAndServe(":8080", nil)
	if errServeWs != nil {
		log.Printf("failed listen 8080")

	} else {
		log.Printf("success listen 8080")
	}
}

func (wss *WebSocketService) wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// helpful log statement to show connections
	log.Println("Client Connected")
	// put in map
	utils.WsMapped = make(map[string]*websocket.Conn)
	utils.WsMapped["default1"] = ws

	wss.ReadRequest(ws)
}

// func handle request from client or udp
func (wss *WebSocketService) ReadRequest(conn *websocket.Conn) {

	go func() {
		for {
			// read in a message
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			// print out that message for clarity
			log.Printf("request from client: %s, msgType: %d", string(p), messageType)
			// pasing to bridge
			wss.Bc.ParseRequestWebsocket(conn.RemoteAddr().String(), string(p))

			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}

		}
	}()

}
