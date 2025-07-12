package tictactoe

import (
	"fmt"
)

type pos [2]int

func (p pos) split() (row int, col int) {
	return p[0], p[1]
}

type Mark rune

func (m Mark) validate() bool {
	if m != MarkX && m != MarkO {
		return false
	}

	return true
}

const (
	MarkNone Mark = 'N'
	MarkX    Mark = 'X'
	MarkO    Mark = 'O'
)

type cells [3][3]Mark

func (c cells) checkPos(pos pos) bool {
	row, col := pos.split()
	if row >= 3 || row < 0 || col >= 3 || col < 0 {
		return false
	}

	return true
}

var win_conditions = [8][3]pos{
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {1, 1}, {2, 2}},
	{{2, 0}, {1, 1}, {0, 2}},
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
}

type Board struct {
	Cells  cells
	Winner Mark
	Next   Mark
}

func NewBoard() Board {
	result := Board{
		Cells: cells{
			{MarkNone, MarkNone, MarkNone},
			{MarkNone, MarkNone, MarkNone},
			{MarkNone, MarkNone, MarkNone},
		},
		Winner: MarkNone,
	}

	return result
}

func (b *Board) Place(m Mark, pos pos) error {
	if !m.validate() {
		return fmt.Errorf("cannot place mark %q", m)
	}

	if m != b.Next {
		return fmt.Errorf("not %q's turn, cannot place", m)
	}

	if !b.Cells.checkPos(pos) {
		return fmt.Errorf("position %v out of bounds", pos)
	}

	if b.checkWinner(m) {
		b.Winner = m
	}

	row, col := pos.split()
	(*b).Cells[row][col] = m
	return nil
}

func (b *Board) checkWinner(m Mark) bool {
	filteredPos := [3][3]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}

	for i, row := range b.Cells {
		for j, c := range row {
			if c == m {
				filteredPos[i][j] = true
			}
		}
	}

sol_loop:
	for _, sol := range win_conditions {
		for _, pos := range sol {
			row, col := pos.split()
			if !filteredPos[row][col] {
				continue sol_loop
			}
		}

		return true
	}

	return false
}
