package board

type Board struct {
	Width  int
	Height int
	Data   [][]rune
}

func New(width int, height int) Board {
	b := Board{Width: width, Height: height}
	b.Data = make([][]rune, b.Height)
	for i := range b.Data {
		b.Data[i] = make([]rune, b.Width)
	}
	return b
}
