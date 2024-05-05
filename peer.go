package main

import (
	"log/slog"
	"net"
)

type Peer struct {
	conn  net.Conn
	msgCh chan []byte
}

func (p *Peer) reedLoop() error {
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(make([]byte, 1024))
		if err != nil {
			slog.Error("read error", "err", err)
			return err
		}
		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])
		p.msgCh <- msgBuf
	}
}

func NewPeer(conn net.Conn, msgCh chan []byte) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
	}
}
