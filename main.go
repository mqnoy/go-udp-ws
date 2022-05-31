package main

import (
	"bebuhcon/utils"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var Udp *net.PacketConn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {

}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	// listen websocket
	go func() {
		setupRoutes()
		errServeWs := http.ListenAndServe(":8080", nil)
		if errServeWs != nil {
			log.Printf("failed listen 8080")

		} else {
			log.Printf("success listen 8080")
		}
	}()

	// listen to incoming udp packets
	Udp, errorUdp := net.ListenPacket("udp", ":9001")
	if errorUdp != nil {
		log.Printf("failed listen 9001")
	} else {
		log.Printf("success listen 9001")
	}

	defer Udp.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := Udp.ReadFrom(buf)

		if err != nil {
			continue
		}
		// go serve(pc, addr, buf[:n])

		log.Println("Received ", string(buf[0:n]), " from ", addr)

		// handle Datagram from bebuh
		HandleDatagram(Udp, addr, string(buf[0:n]))
	}

}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		log.Printf("request from client: %s", string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// helpful log statement to show connections
	log.Println("Client Connected")

	reader(ws)
}

func writeUdp(pc net.PacketConn, addr net.Addr, buf []byte) {
	// 0 - 1: ID
	// 2: QR(1): Opcode(4)
	buf[2] |= 0x80 // Set QR bit

	pc.WriteTo(buf, addr)

}

// func handle datagram
func HandleDatagram(pc net.PacketConn, addr net.Addr, datagram string) {

	datagramParsed := strings.Split(datagram, " ")

	cmdType := datagramParsed[1]
	if cmdType == utils.CmdAck {
		msgType := datagramParsed[5]
		if msgType == utils.SetTime {
			// GetCellPara(pc, addr)
			// StartCell(pc, addr)
		}

	} else {
		// not ack command
		if cmdType == utils.HeartBeat {
			// SetTime(pc, addr)

		}

	}

}

// func setTime
func SetTime(pc net.PacketConn, addr net.Addr) {
	command := []byte("SetTime")
	pc.WriteTo(command, addr)
}

func GetCellPara(pc net.PacketConn, addr net.Addr) {
	command := []byte("GetCellPara")
	pc.WriteTo(command, addr)
}

func StartCell(pc net.PacketConn, addr net.Addr) {
	command := []byte("StartCell")
	pc.WriteTo(command, addr)
}
