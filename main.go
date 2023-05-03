package main

import (
	"fmt"
	"github.com/pinguin-frosch/connect-go/board"
)

func main() {
	fmt.Println("Connect four!")

	b := board.New(6, 7)
	fmt.Println(b.Data)
}
