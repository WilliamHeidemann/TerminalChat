package main

import "strings"

func (s *server) msg(c *client, args []string) {
	if len(args) < 2 {
		c.msg("message is required. usage: /msg MSG")
		return
	}

	msg := strings.Join(args[1:], " ")
	s.broadcast(c, c.nick+": "+msg)
}

func (s *server) broadcast(sender *client, msg string) {
	for addr, m := range s.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}
