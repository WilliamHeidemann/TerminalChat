package main

import "log"

func (s *server) quit(c *client) {
	log.Printf("client has left the chat: %s", c.conn.RemoteAddr().String())

	c.messageClient("sad to see you go :(")
	c.conn.Close()
}
