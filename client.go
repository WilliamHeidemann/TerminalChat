package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn   net.Conn
	nick   string
	server *server
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
			c.server.nick(c, args)
		case "/msg":
			c.server.sendMessage(c, args)
		case "/quit":
			c.server.quit(c)
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("err: " + err.Error() + "\n"))
}

func (c *client) messageClient(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
