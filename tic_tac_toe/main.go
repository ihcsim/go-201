package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	rowSize      = 3
	columnSize   = 3
	nought       = "o"
	cross        = "x"
	noughtStrike = "ooo"
	crossStrike  = "xxx"
)

var (
	starter        = &player{symbol: nought}
	noughtStartPos = []int{0, 0}
	crossStartPos  = []int{1, 2}
)

type board struct {
	m      [][]string
	nought *player
	cross  *player
	active *player
}

func NewBoard() *board {
	b := &board{
		m:      make([][]string, rowSize),
		nought: &player{symbol: nought},
		cross:  &player{symbol: cross},
	}

	for i := 0; i < rowSize; i++ {
		b.m[i] = make([]string, columnSize)
	}

	b.nought.board, b.cross.board = b, b
	return b
}

func main() {
	b := NewBoard()
	b.init(starter, noughtStartPos, crossStartPos)
	for !b.noughtWins() && !b.crossWins() && !b.full() {
		if err := b.switchPlayer(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		x, y, err := b.active.nextCoordinate()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		b.active.move(x, y)
	}

	if b.noughtWins() {
		fmt.Println("Nought wins!")
	} else if b.crossWins() {
		fmt.Println("Cross wins!")
	} else {
		fmt.Println("Tie game!")
	}
}

func (b *board) init(starter *player, noughtPos, crossPos []int) {
	b.active = starter
	b.nought.move(noughtPos[0], noughtPos[1])
	b.cross.move(crossPos[0], crossPos[1])
}

func (b *board) noughtWins() bool {
	for _, row := range b.rows() {
		if row == noughtStrike {
			return true
		}
	}

	for _, column := range b.columns() {
		if column == noughtStrike {
			return true
		}
	}

	for _, diagonal := range b.diagonals() {
		if diagonal == noughtStrike {
			return true
		}
	}

	return false
}

func (b *board) crossWins() bool {
	for _, row := range b.rows() {
		if row == crossStrike {
			return true
		}
	}

	for _, column := range b.columns() {
		if column == crossStrike {
			return true
		}
	}

	for _, diagonal := range b.diagonals() {
		if diagonal == crossStrike {
			return true
		}
	}

	return false
}

func (b *board) switchPlayer() error {
	switch b.active.symbol {
	case nought:
		b.active = b.cross
	case cross:
		b.active = b.nought
	default:
		return fmt.Errorf("Don't know how to handle player with symbol %q", b.active.symbol)
	}

	return nil
}

func (b *board) update(x, y int, symbol string) {
	b.m[x][y] = symbol
}

func (b *board) filled(x, y int) bool {
	return b.m[x][y] != ""
}

func (b *board) full() bool {
	for i := 0; i < rowSize; i++ {
		for j := 0; j < columnSize; j++ {
			if b.m[i][j] == "" {
				return false
			}
		}
	}
	return true
}

type player struct {
	symbol         string
	lastCoordinate []int
	*board
}

func (p *player) move(x, y int) {
	if p.lastCoordinate != nil {
		lastX, lastY := p.lastCoordinate[0], p.lastCoordinate[1]
		fmt.Printf("%s moves from (%d, %d) to (%d, %d)\n", p.symbol, lastX, lastY, x, y)
	}
	p.lastCoordinate = []int{x, y}
	p.update(x, y, p.symbol)
}

func (p *player) nextCoordinate() (x, y int, e error) {
	x, y = p.lastCoordinate[0], p.lastCoordinate[1]
	var availableMoves [][]int
	if x == 0 && y == 0 {
		availableMoves = [][]int{
			[]int{0, 1},
			[]int{1, 1},
			[]int{1, 0},
		}
	} else if x == 0 && y == 1 {
		availableMoves = [][]int{
			[]int{0, 0},
			[]int{0, 2},
			[]int{1, 0},
			[]int{1, 1},
			[]int{1, 2},
		}
	} else if x == 0 && y == 2 {
		availableMoves = [][]int{
			[]int{0, 1},
			[]int{1, 1},
			[]int{1, 2},
		}
	} else if x == 1 && y == 0 {
		availableMoves = [][]int{
			[]int{0, 0},
			[]int{0, 1},
			[]int{1, 1},
			[]int{2, 1},
			[]int{2, 0},
		}
	} else if x == 1 && y == 1 {
		availableMoves = [][]int{
			[]int{0, 0},
			[]int{0, 1},
			[]int{0, 2},
			[]int{1, 0},
			[]int{1, 2},
			[]int{2, 0},
			[]int{2, 1},
			[]int{2, 2},
		}
	} else if x == 1 && y == 2 {
		availableMoves = [][]int{
			[]int{0, 1},
			[]int{0, 2},
			[]int{1, 1},
			[]int{2, 1},
			[]int{2, 2},
		}
	} else if x == 2 && y == 0 {
		availableMoves = [][]int{
			[]int{1, 0},
			[]int{1, 1},
			[]int{2, 1},
		}
	} else if x == 2 && y == 1 {
		availableMoves = [][]int{
			[]int{2, 0},
			[]int{1, 0},
			[]int{1, 1},
			[]int{1, 2},
			[]int{2, 2},
		}
	} else if x == 2 && y == 2 {
		availableMoves = [][]int{
			[]int{2, 1},
			[]int{1, 1},
			[]int{1, 2},
		}
	} else {
		e = fmt.Errorf("Don't know how to handle coordinate %v", p.lastCoordinate)
	}

	for _, move := range availableMoves {
		if p.board.filled(move[0], move[1]) {
			continue
		}
		x, y = move[0], move[1]
	}

	return
}
func (b *board) rows() []string {
	return []string{
		strings.Join(b.m[0], ""),
		strings.Join(b.m[1], ""),
		strings.Join(b.m[2], ""),
	}
}

func (b *board) columns() []string {
	return []string{
		strings.Join([]string{b.m[0][0], b.m[1][0], b.m[2][0]}, ""),
		strings.Join([]string{b.m[0][1], b.m[1][1], b.m[2][1]}, ""),
		strings.Join([]string{b.m[0][2], b.m[1][2], b.m[2][2]}, ""),
	}
}

func (b *board) diagonals() []string {
	return []string{
		strings.Join([]string{b.m[0][0], b.m[1][1], b.m[2][2]}, ""),
		strings.Join([]string{b.m[0][2], b.m[1][1], b.m[2][0]}, ""),
	}
}
