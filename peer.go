package main

import "net"

type Peer struct {
	conn net.Conn
}

func (p *Peer) reedLoop() {
	for {

	}
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
	}
}
