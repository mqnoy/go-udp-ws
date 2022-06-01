package utils

import (
	"net"

	"github.com/gorilla/websocket"
)

// mapped udp connection
var UdpMapped map[string]*net.PacketConn

// mapped websocket connection
var WsMapped map[string]*websocket.Conn
