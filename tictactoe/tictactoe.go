package tictactoe

import "errors"

type Mark uint8

const (
	MarkNone Mark = iota
	MarkX
	MarkO
)

type Board struct {
	cells    [9]Mark
	nextMark Mark
}

func NewBoard() Board {
	result := Board{}
	result.nextMark = MarkX
	return result
}

func (b *Board) updateNextMark() error {
	var updatedMark Mark

	switch b.nextMark {
	case MarkX:
		updatedMark = MarkO
	case MarkO:
		updatedMark = MarkX
	default:
		return errors.New("invalid NextMark Property set")
	}

	b.nextMark = updatedMark
	return nil
}

func (b *Board) PlaceMark(m Mark, p int) error {
	if p >= 9 {
		return errors.New("invalid Mark Position")
	}
	if b.nextMark != m {
		return errors.New("invalid Turn")
	}

	b.cells[p] = m
	b.updateNextMark()
	return nil
}

func (b *Board) getPositions(m Mark) []int {
	p := make([]int, 0, 9)
	c := 0

	for i, v := range b.cells {
		if v == m {
			p = p[:c+1]
			p[c] = i
			c++
		}
	}

	return p
}

func (b *Board) CheckWinner(m Mark) bool {
	/*
		win_conditions := [8][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}
		p := b.getPositions(m)

		for _, cond := range win_conditions {
			//TODO
		}
	*/
	return false
}
