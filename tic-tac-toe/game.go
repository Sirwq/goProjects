package game

import (
	"errors"
	"fmt"
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
	board   Board
	players [2]Player
	turn    int
}

func start() *Game {
	game := new(Game)
	game.players[0].Mark = X
	game.players[1].Mark = O
	game.turn = 0
	return game
}

func getTurn() (x, y int) {
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

func (g *Game) MakeTurn(x, y int, mark Cell) error {
	if g.board[x][y] != EmptyCell {
		return errors.New("В этой точке уже есть фигура.")
	}
	g.turn = 1 - g.turn
	// 1 -> 0 || 0 -> 1 instead of bool that i thougt would work
	g.board[x][y] = g.players[g.turn].Mark
	return nil
}
