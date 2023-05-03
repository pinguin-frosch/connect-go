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

func (b *Board) AddPiece(columnIndex int, symbol string) {
	for i := b.Width - 1; i >= 0; i-- {
		if b.Data[columnIndex][i] == "" {
			b.Data[columnIndex][i] = symbol
			break
		}
	}
}

func (b *Board) IsValidColumn(columnIndex int) bool {
	if columnIndex >= b.Height || columnIndex < 0 {
		return false
	}
	for i := b.Width - 1; i >= 0; i-- {
		if b.Data[columnIndex][i] != "" {
			return false
		}
	}
	return true
}

func (b *Board) PrintBoard() {
	for i := 0; i < b.Width; i++ {
		for j := 0; j < b.Height; j++ {
			v := b.Data[j][i]
			if v == "" {
				fmt.Print("·")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}
