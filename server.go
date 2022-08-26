package main

import (
	"errors"
	"fmt"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

// accepts all incoming commands
func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_USERNAME:
			s.username(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_EXIT:
			s.exit(cmd.client, cmd.args)
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("New client has connected: %s", conn.RemoteAddr().String())

	// initializing client
	c := &client{
		conn: conn,
		username: "anonymous",
		commands: s.commands,
	}

	// reads the command input from client
	c.readInput()

}

func (s *server) username(c *client, args []string) {
	c.username = args[1]
	c.msg(fmt.Sprintf("All right, I will call you %s", c.username))
}

func (s *server) join(c *client, args []string) {
	roomName := args[1]

	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name: roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	r.members[c.conn.RemoteAddr()] = c

	s.exitCurrentRoom(c)

	c.room = r 

	r.broadcast(c, fmt.Sprintf("%s has joined the room!", c.username))
	c.msg(fmt.Sprintf("Yo! Welcome to %s", r.name))
}

func (s *server) listRooms(c *client, args []string) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("Available rooms are: %s", strings.Join(rooms, ", ")))
}

func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("You must join the room first!"))
		return
	}

	c.room.broadcast(c, c.username + ": " + strings.Join(args[1:len(args)], " "))
}

func (s *server) exit(c *client, args []string) {
	log.Printf("Client has disconnected: %s", c.conn.RemoteAddr().String())

	s.exitCurrentRoom(c)

	c.msg("Sad to see you go :(")
	c.conn.Close()
}

func (s *server) exitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the room.", c.username))
	}
}