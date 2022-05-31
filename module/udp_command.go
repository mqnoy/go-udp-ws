package module

import "net"

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
