package main

import "strings"

func (s *server) sendMessage(c *client, args []string) {
	if len(args) < 2 {
		c.messageClient("message is required. usage: /messageClient MSG")
		return
	}

	msg := strings.Join(args[1:], " ")
	s.broadcast(c, c.nick+": "+msg)
}

func (s *server) broadcast(sender *client, msg string) {
	for addr, m := range s.members {
		if sender.conn.RemoteAddr() != addr {
			m.messageClient(msg)
		}
	}
}
