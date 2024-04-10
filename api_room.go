package main

import (
	"net/http"
)

type RoomApi struct {
	Name string
	LoggedPlayer string
	OpponentPlayer string
	Turn bool
}

func roomApiHandler(w http.ResponseWriter, r *http.Request) {
}
