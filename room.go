package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pranavtaysheti/goTTT/tictactoe"
)

type Room struct {
	board     tictactoe.Board
	players   [2]*Player
	noPlayers int
}

func NewRoom(name string) Room {
	newRoom := Room{
		board:     tictactoe.NewBoard(),
		noPlayers: 0,
	}
	rooms[name] = &newRoom
	return newRoom
}

func (r *Room) AddPlayer(player *Player) error {
	if r.noPlayers >= 2 {
		return errors.New("room housefull")
	}

	r.players[r.noPlayers] = player
	r.noPlayers++

	//TODO: Send WS signal to notify client to update player list.
	return nil
}

func (r *Room) MarkCell(cell int, playerUuid string) error {
	var playerNo int

	for i := 0; i < 2; i++ {
		if r.players[i].String() == playerUuid {
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

func boardApiHandler(w http.ResponseWriter, r *http.Request) {
	var p tictactoe.Board

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		//TODO: Handle error properly

	}
}

func playerApiHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
}
