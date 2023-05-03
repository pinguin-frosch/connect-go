package main

import (
	"fmt"

	"github.com/pinguin-frosch/connect-go/board"
	"github.com/pinguin-frosch/connect-go/game"
)

func main() {
	fmt.Println("Connect four!")

	b := board.New(6, 7)
	g := game.New(&b)
	g.Setup()
	g.Play()
}
