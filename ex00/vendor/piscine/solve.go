package piscine

import (
	"fmt"
	"ft"
	"os"
)

type Tetrimino struct {
	data []string
}

// https://pkg.go.dev/builtin
// append, cap, copy, delete, len, make, new

// https://pkg.go.dev/os
// os
func Ft_solve(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		// make error the same everywhere
		fmt.Println("Error readfile")
		//remove me
		fmt.Println(err)
		return
	}

	// do we need check if there is errors during parse once more ?
	rawTets := parseData(data)
	fmt.Println("found ", len(rawTets), "tetriminos")

	unnamedTets := []Tetrimino{}
	for _, t := range rawTets {
		unnamedTets = append(unnamedTets, minimizeTet(t))
	}

	//for _, t := range unnamedTets {
	//PrintTetrimino(t)
	//fmt.Println("~~~~")
	//}

	// use A, B, C, ... etc instead of #
	// max # of tetriminos is 26, so we are safe

	tets := []Tetrimino{}
	letter := 'A'
	for _, t := range unnamedTets {
		tets = append(tets, nameTet(t, letter))
		//PrintTetrimino(tets[i])
		letter++
	}

	for _, t := range tets {
		PrintTetrimino(t)
		fmt.Println("~~~~")
	}

	emptyBoard := [][]byte{}
	for i := 0; i < 4; i++ {
		emptyBoard = append(emptyBoard, []byte{})
		for j := 0; j < 4; j++ {
			emptyBoard[i] = append(emptyBoard[i], '.')
		}
	}

	printBoard(emptyBoard)
	fmt.Println("~~~")

	//emptyBoard = boardPlusPlus(emptyBoard)

	//printBoard(emptyBoard)
	//fmt.Println("~~~")

	//
	// increase board until it solved
	for {
		if solve(emptyBoard, tets, 0) {
			break
		}
		emptyBoard = boardPlusPlus(emptyBoard)
	}
	//

	// create a smallest board: 2x2
	// try to put first one here, second, etc until can put something
	// to put element onto board use function
	// keep map with elements already used so we can skip while iterationg over array

	// board is a slice
	// make function to make board bigger
	// make function to check if board if full?
	// iterate over all and try to put into board
	// trying to put into board function (try it over all possible squares on current board

	// take a piece try to put on each point of the board in a loop
	// take second piece and try to put it on each square in a loop
	// ..
	// recursion with deept == len(pieces/Tetrimino)
	// board is copied and passed to inner call
	// when we fit all - we solved! it must be gready solution
}

func solve(b [][]byte, tets []Tetrimino, idx int) bool {
	// solved
	if idx >= len(tets) {
		fmt.Println("solved!!!")
		printBoard(b)
		return true
	}

	t := tets[idx]
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			nb, err := putOnBrd(b, t, i, j)
			if err == false {
				//fmt.Println(string('A' + idx ) + ": able to put at", i, j)
				//printBoard(nb)
				if solve(nb, tets, idx+1) {
					return true
				}
			} else {
				//fmt.Println(string('A' + idx ) + ": unable at", i, j)
			}
		}
	}
	return false
}

func putOnBrd(b [][]byte, t Tetrimino, ii, jj int) ([][]byte, bool) {
	nb := [][]byte{}
	for i := 0; i < len(b); i++ {
		nb = append(nb, []byte{})
		for j := 0; j < len(b[i]); j++ {
			nb[i] = append(nb[i], b[i][j])
		}
	}

	// tet's height overfows
	if ii+len(t.data) > len(b) {
		return nb, true
	}

	// tet's width overflows
	if jj+len(t.data[0]) > len(b[0]) {
		return nb, true
	}

	for i := 0; i < len(t.data); i++ {
		for j := 0; j < len(t.data[i]); j++ {
			// we can put any tet part here
			if nb[ii+i][jj+j] == '.' {
				nb[ii+i][jj+j] = t.data[i][j]
			} else if nb[ii+i][jj+j] != '.' && t.data[i][j] == '.' {
				// noop //
				// O  O //
				//  \\\ //
				//nb[ii+i][jj+j] = t.data[i][j]
			} else {
				return nb, true
			}
		}
	}

	return nb, false
}

func parseData(data []byte) []Tetrimino {
	tets := []Tetrimino{}
	var t Tetrimino
	line := ""
	var prev byte
	prev = 46
	for _, b := range data {
		if b == 10 {
			if prev == 10 {
				tets = append(tets, t)
				t = Tetrimino{}
			} else {
				if !isEmptyLine(line) {
					t.data = append(t.data, line)
				}
				line = ""
			}
		} else {
			line = line + string(b)
		}
		prev = b
	}
	tets = append(tets, t)
	return tets
}

func PrintTetrimino(tet Tetrimino) {
	for _, s := range tet.data {
		fmt.Println(s)
	}
}

func minimizeTet(t Tetrimino) Tetrimino {
	tet := Tetrimino{}
	colsToRemove := [4]int{}
	for j := 0; j < len(t.data[0]); j++ {
		canRemove := true
		for i := 0; i < len(t.data); i++ {
			if t.data[i][j] != 46 {
				canRemove = false
				break
			}
		}
		if canRemove {
			colsToRemove[j] = 1
		}
	}

	for i := 0; i < len(t.data); i++ {
		line := ""
		for j := 0; j < len(t.data[0]); j++ {
			if colsToRemove[j] == 1 {
				continue
			}
			line = line + string(t.data[i][j])
		}
		tet.data = append(tet.data, line)
	}
	return tet
}

func isEmptyLine(s string) bool {
	for _, v := range s {
		if v != 46 {
			return false
		}
	}
	return true
}

func nameTet(t Tetrimino, letter rune) Tetrimino {
	newTet := Tetrimino{}
	for i := 0; i < len(t.data); i++ {
		line := ""
		for j := 0; j < len(t.data[0]); j++ {
			if t.data[i][j] == '#' {
				line = line + string(letter)
			} else {
				line = line + string(t.data[i][j])
			}
		}
		newTet.data = append(newTet.data, line)
	}
	return newTet
}

func printBoard(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			ft.PrintRune(rune(board[i][j]))
		}
		ft.PrintRune('\n')
	}
}

func boardPlusPlus(board [][]byte) [][]byte {
	res := [][]byte{}
	for i := 0; i <= len(board); i++ {
		res = append(res, []byte{})
		for j := 0; j <= len(board[0]); j++ {
			res[i] = append(res[i], '.')
		}
	}
	return res
}
