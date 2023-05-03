package game

import (
	"bufio"
	"fmt"
	"os"

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

func (g *Game) Play() {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; true; i = (i + 1) % 2 {
		g.Board.PrintBoard()

		fmt.Printf("%s, select a column: ", g.Players[i].Symbol)
		var column int

		for {
			_, err := fmt.Scanf("%d", &column)
			if err != nil {
				fmt.Print("Invalid column, try again: ")
				scanner.Scan()
				continue
			}
			if g.Board.IsValidColumn(column) {
				break
			}
			fmt.Print("Invalid column, try again: ")
		}

		fmt.Printf("Adding your piece in column %d\n", column)
		col, row := g.Board.AddPiece(column, g.Players[i].Symbol)

		if g.Board.CheckWin(col, row, g.Players[i].Symbol) {
			fmt.Printf("%s wins!", g.Players[i].Name)
		}
	}
}

func (g *Game) Setup() {
	g.Players[0].Name = "Player 1"
	g.Players[0].Symbol = "C"
	g.Players[1].Name = "Player 2"
	g.Players[1].Symbol = "F"
}
