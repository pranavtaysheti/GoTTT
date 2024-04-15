package server

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pranavtaysheti/GoTTT/tictactoe"
)

type Room struct {
	board     tictactoe.Board
	players   [2]*Player
	noPlayers int
	clients   []*Client
}

func GetOrMakeRoom(name string) *Room {
	for n, r := range Rooms {
		if n == name {
			return r
		}
	}

	return NewRoom(name)
}

func NewRoom(name string) *Room {
	newRoom := &Room{
		board:     tictactoe.NewBoard(),
		noPlayers: 0,
	}
	Rooms[name] = newRoom
	return newRoom
}

func (r *Room) PlayerExists(np *Player) bool {
	for _, p := range r.players {
		if p == np {
			return true
		}
	}

	return false
}

func (r *Room) NotifyClients(m WsMessage) error {
	for _, c := range r.clients {
		if c.socket != nil {
			return c.SendMessage(m)
		}
	}

	return nil
}

func (r *Room) AddPlayer(p *Player) error {
	if r.PlayerExists(p) {
		return nil
	}

	if r.noPlayers >= 2 {
		return errors.New("room housefull")
	}

	r.players[r.noPlayers] = p
	r.noPlayers++

	return nil
}

func (r *Room) AddClient(c *Client) {
	r.clients = append(r.clients, c)
}

func (r *Room) MarkCell(cell int, playerUuid uuid.UUID) error {
	var playerNo int

	for i := 0; i < 2; i++ {
		if r.players[i].uuid == playerUuid {
			playerNo = i
			break
		}
	}

	err := r.board.PlaceMark(tictactoe.Mark(uint8(playerNo)), cell)
	if err != nil {
		return err
	}

	return nil
}

func (r *Room) PlayerNames() []string {
	res := []string{}
	for _, p := range r.players {
		res = append(res, p.name)
	}
	return res
}
