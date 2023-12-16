package main

import (
	"errors"

	"github.com/pranavtaysheti/goTTT/tictactoe"
)

type Room struct {
	name      string
	board     tictactoe.Board
	players   [2]Player
	noPlayers int
}

func NewRoom(name string) Room {
	return Room{
		name:      name,
		board:     tictactoe.NewBoard(),
		noPlayers: 0,
	}
}

func (r *Room) AddPlayer(playerName string) (Player, error) {
	newPlayer := NewPlayer(playerName)

	if r.noPlayers < 2 {
		r.players[r.noPlayers] = newPlayer
		r.noPlayers++
	} else {
		return newPlayer, errors.New("room HouseFull")
	}

	return newPlayer, nil
}

func (r *Room) MarkCell(cell int, playerUuid string) error {
	var playerNo int

	for i := 0; i < 2; i++ {
		if r.players[i].uuid.String() == playerUuid {
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
