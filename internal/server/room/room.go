package room

import "github.com/pranavtaysheti/GoTTT/tictactoe"

type Room struct {
	board     tictactoe.Board
	players   [2]*Player
	noPlayers int
	clients   []*Client
}
