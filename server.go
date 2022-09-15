package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	commands chan command
	members  map[net.Addr]*client
}

func newServer() *server {
	return &server{
		commands: make(chan command),
		members:  make(map[net.Addr]*client),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) *client {
	log.Printf("new client has joined: %s", conn.RemoteAddr().String())
	c := &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}
	s.members[conn.RemoteAddr()] = c

	c.msg("Welcome to the server!")
	c.msg("Available commands are /nick /msg and /quit")

	return c
}

func (s *server) nick(c *client, args []string) {
	if len(args) < 2 {
		c.msg("nick is required. usage: /nick NAME")
		return
	}

	c.nick = args[1]
	c.msg(fmt.Sprintf("Nickname changed to %s", c.nick))
}

func (s *server) msg(c *client, args []string) {
	if len(args) < 2 {
		c.msg("message is required, usage: /msg MSG")
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

func (s *server) quit(c *client) {
	log.Printf("client has left the chat: %s", c.conn.RemoteAddr().String())

	c.msg("sad to see you go =(")
	c.conn.Close()
}
