package main

import "fmt"

func (s *server) nick(c *client, args []string) {
	if len(args) < 2 {
		c.messageClient("nick is required. usage: /nick NAME")
		return
	}

	c.nick = args[1]
	c.messageClient(fmt.Sprintf("Nickname changed to %s", c.nick))
}
