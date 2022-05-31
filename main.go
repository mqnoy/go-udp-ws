package main

import (
	us "bebuhcon/services/udp_services"
	ws "bebuhcon/services/websocket_services"
)

func init() {

}

func main() {
	// datagram handle
	Hd := &us.HandleDatagram{}

	// listen websocket
	Wss := &ws.WebSocketService{}
	go func() {
		Wss.Start()
	}()

	// listen to incoming udp
	Uss := &us.UdpSocketService{Hd: Hd}
	Uss.Start()

}
