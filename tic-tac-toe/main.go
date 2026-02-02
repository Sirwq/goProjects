package main

import (
	"fmt"
	"tic-tac-toe/game"
)

func main() {
	g := game.New()
	winner := g.CheckWin()

	for !g.IsOver() {
		x, y := game.GetTurn()
		for {
			if g.MakeTurn(x, y) == nil {
				break
			}
			x, y = game.GetTurn()
		}
		fmt.Println(g)
		winner = g.CheckWin()
		if winner != game.EmptyCell {
			break
		}
	}
	winner = g.CheckWin()
	switch winner {
	case game.X:
		fmt.Println("X is winner")
	case game.O:
		fmt.Println("O is winner")
	default:
		fmt.Println("it's a tie")
	}
}
