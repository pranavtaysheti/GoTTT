package main

import (
	"errors"
	"github.com/google/uuid"
)

type Player uuid.UUID

func NewPlayer(name string) *Player {
	player := Player(uuid.New())
	players[name] = &player
	return &player
}

func (p Player) String() string {
	return uuid.UUID(p).String()
}

func getPlayerbyUUID(uuid string) (*Player, error) {
	for _, p := range players {
		if p.String() == uuid {
			return p, nil
		}
	}

	return &Player{}, errors.New("Player of given UUID not found")
}

