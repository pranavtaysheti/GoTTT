package main

import (
	"github.com/google/uuid"
)

type Player uuid.UUID

var players = make(map[string]*Player, 100)

func NewPlayer(name string) *Player {
	player := Player(uuid.New())
	players[name] = &player
	return &player
}

func (p Player) String() string {
	return uuid.UUID(p).String()
}

func getOrMakePlayer(name string) *Player {
	for n, p := range players {
		if n == name {
			return p
		}
	}

	return NewPlayer(name)
}
