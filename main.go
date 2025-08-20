package main

import (
	"fmt"
	"math/rand"
)

// import "time"

type Board struct {
	cells [Size][Size]int
}

func (b Board) Print() {
	hline := "+------+------+------+------+\n"
	fmt.Print(hline)

	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if b.cells[r][c] == 0 {
				fmt.Printf("| %4s ", "")
			} else {
				fmt.Printf("| %4d ", b.cells[r][c])
			}
		}
		fmt.Println("|")
		fmt.Print(hline)
	}
}

func (b *Board) AddRandomTile() {
	var empty []struct{ r, c int }
	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if b.cells[r][c] == 0 {
				empty = append(empty, struct{ r, c int }{r, c})
			}
		}
	}

	if len(empty) == 0 {
		return
	}

	choice := empty[rand.Intn(len(empty))]

	// place a 2 (90% chance) or 4 (10% chance)
	if rand.Intn(10) == 0 {
		b.cells[choice.r][choice.c] = 4
	} else {
		b.cells[choice.r][choice.c] = 2
	}
}

func (b *Board) MoveLeft() {
	for r := 0; r < Size; r++ { // Changed from Size - 1 to Size
		slice := mergeRowLeft(b.cells[r][:])
		// Convert slice back to array
		for i := 0; i < Size; i++ {
			b.cells[r][i] = slice[i]
		}
	}
}

func (b *Board) MoveRight() {
	for r := range Size {
		reversed := reverseRow((b.cells[r][:]))
		mergeRow := mergeRowLeft(reversed)
		finalReversed := reverseRow(mergeRow)
		// Convert slice back to array
		for i := range Size {
			b.cells[r][i] = finalReversed[i]
		}
	}

}

func (b *Board) Transpose() {
	for r := range b.cells {
		for c := range b.cells[r] {
			b.cells[r][c], b.cells[c][r] = b.cells[c][r], b.cells[r][c]
		}
	}
}

func (b *Board) MoveUp() {
	b.Transpose()
	b.MoveLeft()
	b.Transpose()
}

func (b *Board) MoveDown() {
	b.Transpose()
	b.MoveRight()
	b.Transpose()
}

func main() {
	var b Board
	b.cells[0][1] = 16
	b.cells[0][2] = 16
	b.AddRandomTile()
	b.AddRandomTile()
	b.Print()

	b.MoveLeft()
	b.Print()

	// b.MoveRight()
	// b.Print()

}
