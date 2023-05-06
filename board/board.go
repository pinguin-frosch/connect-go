package board

import (
	"fmt"
)

type Board struct {
	Width  int
	Height int
	Data   [][]string
}

func New(width int, height int) Board {
	b := Board{Width: width, Height: height}
	b.Data = make([][]string, b.Height)
	for i := range b.Data {
		b.Data[i] = make([]string, b.Width)
	}
	return b
}

func (b *Board) AddPiece(columnIndex int, symbol string) (column int, row int) {
	for i := b.Width - 1; i >= 0; i-- {
		if b.Data[columnIndex][i] == "" {
			b.Data[columnIndex][i] = symbol
			return columnIndex, i
		}
	}
	return 0, 0
}

// FIXME: Should also check for draw and also checks win states with the piece
// in the middle
func (b *Board) CheckWin(column int, row int, symbol string) bool {
	counts := map[string]int{}

	counts["LD"] = 1
	for x := 1; column-x >= 0 && row-x >= 0; x++ {
		if b.Data[column-x][row-x] != symbol {
			break
		}
		counts["LD"]++
	}
	for x := 1; column+x < b.Height && row+x < b.Width; x++ {
		if b.Data[column+x][row+x] != symbol {
			break
		}
		counts["LD"]++
	}

	counts["V"] = 1
	for x := 1; column-x >= 0; x++ {
		if b.Data[column-x][row] != symbol {
			break
		}
		counts["V"]++
	}
	for x := 1; column+x < b.Height; x++ {
		if b.Data[column+x][row] != symbol {
			break
		}
		counts["V"]++
	}

	counts["RD"] = 1
	for x := 1; column-x >= 0 && row+x < b.Width; x++ {
		if b.Data[column-x][row+x] != symbol {
			break
		}
		counts["RD"]++
	}
	for x := 1; column+x < b.Height && row-x >= 0; x++ {
		if b.Data[column+x][row-x] != symbol {
			break
		}
		counts["RD"]++
	}

	counts["H"] = 1
	for x := 1; row+x < b.Width; x++ {
		if b.Data[column][row+x] != symbol {
			break
		}
		counts["H"]++
	}

	for _, v := range counts {
		if v >= 4 {
			return true
		}
	}

	return false
}

func (b *Board) IsValidColumn(columnIndex int) bool {
	if columnIndex >= b.Height || columnIndex < 0 {
		return false
	}
	used := 0
	for i := b.Width - 1; i >= 0; i-- {
		if b.Data[columnIndex][i] != "" {
			used++
		}
	}
	if used < b.Width {
		return true
	} else {
		return false
	}
}

func (b *Board) PrintBoard() {
	fmt.Println()
	for i := 0; i < b.Height; i++ {
		fmt.Printf(" %d", i+1)
	}
	fmt.Println()

	for i := 0; i < b.Width; i++ {
		for j := 0; j < b.Height; j++ {
			v := b.Data[j][i]
			fmt.Print("|")
			if v == "" {
				fmt.Print("·")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Print("|\n")
	}

	fmt.Print("+")
	for i := 0; i < b.Height-1; i++ {
		fmt.Print("--")
	}
	fmt.Print("-+\n\n")
}
