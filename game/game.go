package game

import (
	"github.com/pinguin-frosch/connect-go/board"
	"github.com/pinguin-frosch/connect-go/player"
)

type Game struct {
	Players [2]player.Player
	Board   *board.Board
}

func New(board *board.Board) Game {
	g := Game{Board: board}
	return g
}

func (g *Game) Setup() {
	g.Players[0].Name = "Player 1"
	g.Players[0].Symbol = 'C'
	g.Players[1].Name = "Player 2"
	g.Players[1].Symbol = 'F'
}
