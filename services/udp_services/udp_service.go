package services

import (
	bc "bebuhcon/controller"
	"bebuhcon/utils"
	"log"
	"net"
)

type UdpSocketService struct {
	Bc *bc.BridgeController
}

var Pc net.PacketConn

// func start open port udp
func (uss *UdpSocketService) Start() {
	log.Printf("UdpSocketService::Start executed")

	// listen to incoming udp packets
	pc, errorUdp := net.ListenPacket("udp", ":9001")

	if errorUdp != nil {
		log.Printf("failed listen 9001")
	} else {
		log.Printf("success listen 9001")

		// put in map
		utils.UdpMapped = make(map[string]*net.PacketConn)
		utils.UdpMapped["default1"] = &pc
		uss.ReadRequest(pc)

	}

}

// func handle request from client or ws
func (uss *UdpSocketService) ReadRequest(pc net.PacketConn) {
	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)

		if err != nil {
			continue
		}
		// go serve(pc, addr, buf[:n])

		log.Println("UdpSocketService::ReadRequest - Received ", string(buf[0:n]), " from ", addr)

		// handle Datagram from bebuh
		// uss.Hd.Parsing(pc, addr.S, string(buf[0:n]))
		uss.Bc.ParseDatagramUdp(addr.String(), string(buf[0:n]))
	}
}

// func write datagram over udp
func (wss *UdpSocketService) WriteDatagram(addr net.Addr, command string) {
	p := []byte(command)
	Pc.WriteTo(p, addr)
}
