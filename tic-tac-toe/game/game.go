package game

import (
	"errors"
	"fmt"
	"strings"
)

type Cell int8

const (
	EmptyCell = iota
	X
	O
)

type Board [3][3]Cell

type Player struct {
	Mark Cell
}

type Game struct {
	Board   Board // by convention public fields stats with big letter
	Players [2]Player
	turn    int
}

func New() *Game {
	return &Game{
		Board:   Board{},
		Players: [2]Player{Player{X}, Player{O}},
		turn:    0}
}

func GetTurn() (x, y int) {
	for {
		fmt.Println("Введите координаты, куда хотите поставить вашу фигуру:")

		n, err := fmt.Scan(&x, &y)

		if err != nil || n != 2 {
			fmt.Println("Введены некоректные данные")
			fmt.Scanln() // c-like buffer clear
			continue
		}

		if x < 0 || x > 2 || y < 0 || y > 2 {
			fmt.Println("Координаты должны соответствовать сетке")
			continue
		}
		break
	}
	return
}

func (g *Game) MakeTurn(x, y int) error {
	if g.Board[x][y] != EmptyCell {
		return errors.New("В этой точке уже есть фигура.")
	}
	g.Board[x][y] = g.Players[g.turn].Mark
	g.turn = 1 - g.turn
	// 1 -> 0 || 0 -> 1 instead of bool that i thougt would work
	return nil
}

func (g Game) CheckWin() Cell { // ILL REFACTOR THIS FOR SUUUUUURE but not RN
	// Diagonal
	if g.Board[0][0] != EmptyCell &&
		g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] {
		return g.Board[1][1]
	}
	// Reverse diagonal
	if g.Board[2][0] != EmptyCell &&
		g.Board[2][0] == g.Board[1][1] && g.Board[1][1] == g.Board[0][2] {
		return g.Board[1][1]
	}

	if g.Board[0][0] != EmptyCell &&
		g.Board[0][0] == g.Board[0][1] && g.Board[0][1] == g.Board[0][2] {
		return g.Board[0][0]
	}

	if g.Board[1][0] != EmptyCell &&
		g.Board[1][0] == g.Board[1][1] && g.Board[1][1] == g.Board[1][2] {
		return g.Board[1][1]
	}

	if g.Board[2][0] != EmptyCell &&
		g.Board[2][0] == g.Board[2][1] && g.Board[2][1] == g.Board[2][2] {
		return g.Board[2][2]
	}

	if g.Board[0][0] != EmptyCell &&
		g.Board[0][0] == g.Board[0][1] && g.Board[0][1] == g.Board[0][2] {
		return g.Board[0][0]
	}
	if g.Board[0][1] != EmptyCell &&
		g.Board[0][1] == g.Board[1][1] && g.Board[1][1] == g.Board[2][1] {
		return g.Board[1][1]
	}
	if g.Board[0][2] != EmptyCell &&
		g.Board[0][2] == g.Board[1][2] && g.Board[1][2] == g.Board[2][2] {
		return g.Board[2][2]
	}

	return EmptyCell
}

func (g Game) IsOver() bool {
	for i := range g.Board {
		for j := range g.Board[0] {
			if g.Board[i][j] == EmptyCell {
				return false
			}
		}
	}
	return true
}

func (c Cell) String() string {
	switch c {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return " "
	}
}

func (g Game) String() string {
	var sb strings.Builder

	for i := range g.Board {
		for j := range g.Board[0] {
			sb.WriteString(g.Board[i][j].String())

			if j < 2 {
				sb.WriteString("|")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
