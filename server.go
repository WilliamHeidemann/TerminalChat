package main

import (
	"log"
	"net"
)

type server struct {
	members map[net.Addr]*client
}

func newServer() *server {
	return &server{
		members: make(map[net.Addr]*client),
	}
}

func (s *server) newClient(conn net.Conn) *client {
	log.Printf("new client has joined: %s", conn.RemoteAddr().String())
	c := &client{
		conn:   conn,
		nick:   "anonymous",
		server: s,
	}
	s.members[conn.RemoteAddr()] = c

	c.messageClient("Welcome to the server!")
	c.messageClient("Available commands are /nick /messageClient and /quit")

	return c
}
