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

// TODO: implement this properly
func (b *Board) CheckWin(column int, row int, symbol string) bool {
	counts := map[string]int{}
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
