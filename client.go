package main

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type client struct {
	room   *Room
	socket *websocket.Conn
}

var connections = make(map[uuid.UUID]*client, 100)

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
	return nil
}

func (c *client) AddPlayer(p *Player) error {
	if c.room == nil {
		return errors.New("Cant add player, room not associated with client")
	}

	return c.room.AddPlayer(p)
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

	return nil
}