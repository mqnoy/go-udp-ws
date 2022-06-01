package controller

import (
	"bebuhcon/utils"
	"log"
	"net"
)

// us "bebuhcon/services/udp_services"
// ws "bebuhcon/services/websocket_services"

type BridgeController struct {
	// Wss *ws.WebSocketService
	// Uss *us.UdpSocketService
}

func (b *BridgeController) Handle() {

}

// func parsing datagram udp
func (uss *BridgeController) ParseDatagramUdp(addr string, datagram string) {

	log.Println("BridgeController::ReadRequest - Received ", datagram, " from ", addr)

	wn := utils.WsMapped["default1"]

	p := []byte("command")

	if err := wn.WriteMessage(1, p); err != nil {
		log.Println(err)
	}

}

// func parsing request websocket
func (uss *BridgeController) ParseRequestWebsocket(addr string, requests string) {
	log.Println("BridgeController::ParseRequestWebsocket - Received ", requests, " from ", addr)

	var np *net.PacketConn = utils.UdpMapped["default1"]
	var pc net.PacketConn = *np

	p := []byte("command")
	ipaddr, err := net.ResolveIPAddr("ip", addr)
	if err != nil {
		// handle error
	} else {

		if _, errWu := pc.WriteTo(p, ipaddr); errWu != nil {
			log.Println(errWu)
		}
	}
}
