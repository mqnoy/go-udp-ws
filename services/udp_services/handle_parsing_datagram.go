package services

import (
	"bebuhcon/utils"
	"net"
	"strings"
)

type HandleDatagram struct {
}

// func handle datagram
func (hd *HandleDatagram) Parsing(pc net.PacketConn, addr net.Addr, datagram string) {

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
			// TODO: write command from handle parsing datagram
			p := []byte("SetTime")
			pc.WriteTo(p, addr)

		}

	}

}
