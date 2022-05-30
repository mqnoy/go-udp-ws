package main

import (
	"bebuhcon/utils"
	"fmt"
	"log"
	"net"
	"strings"
)

var Udp *net.PacketConn

func init() {

}

func main() {
	// listen to incoming udp packets
	Udp, errorUdp := net.ListenPacket("udp", ":9001")
	if errorUdp != nil {
		log.Fatal(errorUdp)
	}

	defer Udp.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := Udp.ReadFrom(buf)

		if err != nil {
			continue
		}
		// go serve(pc, addr, buf[:n])

		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		// handle Datagram from bebuh
		HandleDatagram(Udp, addr, string(buf[0:n]))
	}

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
		msgType := datagramParsed[5];
		if msgType == utils.SetTime{
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

func GetCellPara(pc net.PacketConn, addr net.Addr){
	command := []byte("GetCellPara")
	pc.WriteTo(command, addr)
}

func StartCell(pc net.PacketConn, addr net.Addr){
	command := []byte("StartCell")
	pc.WriteTo(command, addr)
}
