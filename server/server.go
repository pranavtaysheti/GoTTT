package server

import (
	"errors"
	"github.com/google/uuid"
)

var Connections = make(map[uuid.UUID]*Client, 100)
var Rooms = make(map[string]*Room, 100)

func GetClientByUUIDString(u string) (*Client, error) {
	for id, c := range Connections {
		if id.String() == u {
			return c, nil
		}
	}

	return &Client{}, errors.New("Client with given uuid not found")
}
