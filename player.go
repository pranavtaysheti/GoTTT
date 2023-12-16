package main

import "github.com/google/uuid"

type Player struct {
	name string
	uuid uuid.UUID
}

func NewPlayer(n string) Player {
	return Player{
		name: n,
		uuid: uuid.New(),
	}
}
