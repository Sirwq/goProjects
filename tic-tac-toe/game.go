package game

import "fmt"

type Cell int8

const (
	EmptyCell = iota
	X
	O
)

type Board [3][3]Cell

type Player Cell

type Game struct {
	board   Board
	players [2]Player
	curturn bool
}

func start() *Game {
	game := new(Game)
	game.players[0] = X
	game.players[1] = O
	game.curturn = false
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

}
