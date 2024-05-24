package server

import (
	"github.com/google/uuid"
)

type Player struct {
	uuid uuid.UUID
	name  string
}

func NewPlayer(name string) *Player {
	player := Player{
		uuid: uuid.New(),
		name: name,
	}

	return &player
}

func (p *Player) GetName() string {
	return p.name 
}
