package main

import (
	"log"
	"net"
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
