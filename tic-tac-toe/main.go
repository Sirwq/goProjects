package main

import (
	"fmt"
	"tic-tac-toe/game"
)

func main() {
	g := game.New()
	g.MakeTurn(1, 1, game.X)

	fmt.Println(g)
}
