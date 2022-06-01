package main

import (
	"bebuhcon/controller"
	us "bebuhcon/services/udp_services"
	ws "bebuhcon/services/websocket_services"
)

func init() {

}

func main() {
	// datagram handle
	Bc := &controller.BridgeController{}

	// listen websocket
	Wss := &ws.WebSocketService{Bc: Bc}
	go func() {
		Wss.Start()
	}()

	// listen to incoming udp
	Uss := &us.UdpSocketService{Bc: Bc}
	Uss.Start()

}
