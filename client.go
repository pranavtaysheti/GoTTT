package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type client struct {
	player *Player
	room   *Room
	socket *websocket.Conn
}

var connections = make(map[uuid.UUID]*client, 100)

func getClientByUUIDString(u string) (*client, error) {
	for id, c := range connections {
		if id.String() == u {
			return c, nil
		}
	}

	return &client{}, errors.New("Client with given uuid not found")
}

func NewClient() uuid.UUID {
	u := uuid.New()
	connections[u] = &client{}
	return u
}

func (c *client) SetRoom(r *Room) error {
	if c.room != nil {
		return errors.New("Client already associated with a room")
	}

	c.room = r
	c.room.AddClient(c)
	return nil
}

func (c *client) AddPlayer(p *Player) error {
	if c.room == nil {
		return errors.New("Cant add player, room not associated with client")
	}

	return c.room.AddPlayer(p)
}

func (c *client) SetPlayer(p *Player) error {
	if c.player != nil {
		return errors.New("Client already associated with player")
	}

	c.player = p
	return nil
}

func (c *client) SendMessage (m WsMessage) error {
	if c.socket == nil {
		return errors.New("websocket not initialized")
	}

	_, err := c.socket.Write([]byte(m))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
func (c *client) Login (p *Player, r *Room) error {
	err := c.SetRoom(r)
	if err != nil {
		return err
	}

	err = c.AddPlayer(p)
	if err != nil {
		return err
	}

	err = c.SetPlayer(p)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) ConnectSocket (ws *websocket.Conn) error {
	if c.socket != nil {
		return errors.New("Client already associated with socket")
	}

	c.socket = ws
	return nil
}

func getClientFromRequest(r *http.Request) (*client, error) {
	c, err := getClientCookie(r)
	if err != nil {
		return nil, err 
	}

	cl, err := getClientFromCookie(c)
	if err != nil {
		return nil, err 
	}

	return cl, nil
}

