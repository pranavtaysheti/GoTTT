package server

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type Client struct {
	player *Player
	room   *Room
	socket *websocket.Conn
}


func NewClient() uuid.UUID {
	u := uuid.New()
	Connections[u] = &Client{}
	return u
}


func (c *Client) GetRoom() (*Room, error) {
	if c.room == nil {
		return nil, errors.New("Room not set on this client")
	}

	return c.room, nil
}

func (c *Client) SetRoom(r *Room) error {
	if c.room != nil {
		return errors.New("room already associated with client")
	}

	c.room = r
	c.room.AddClient(c)
	return nil
}

func (c *Client) UnsetRoom(r *Room) error {
	if c.room == nil {
		return errors.New("Room already unset.")
	}

	c.room = nil
	return nil
}

func (c *Client) GetPlayer() (*Player, error) {
	if c.player == nil {
		return nil, errors.New("Player not associated with client")
	}

	return c.player, nil
}

func (c *Client) AddPlayer(p *Player) error {
	room, err := c.GetRoom()
	if err != nil {
		return err
	}

	return room.AddPlayer(p)
}

func (c *Client) SetPlayer(p *Player) error {
	if c.player != nil {
		return errors.New("Client already associated with player")
	}

	c.player = p
	return nil
}

func (c *Client) SendMessage (m WsMessage) error {
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
func (c *Client) Login (p *Player, r *Room) error {
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

func (c *Client) ConnectSocket (ws *websocket.Conn) error {
	if c.socket != nil {
		return errors.New("Client already associated with socket")
	}

	c.socket = ws
	return nil
}

func (c *Client) DisconnectSocket() error {
	if c.socket == nil {
		return errors.New("Socket already unset.")
	}

	c.socket = nil
	return nil
}
