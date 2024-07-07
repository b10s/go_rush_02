package piscine

import (
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
		Ft_PrintErr()
		return
	}

	rawTets, ok := parseData(data)
	if !ok {
		Ft_PrintErr()
		return
	}
	// impossible for this algorithm
	if len(rawTets) >26 {
		Ft_PrintErr()
		return
	}

	cleanedTets := []Tetrimino{}
	for _, t := range rawTets {
		cleanedTets = append(cleanedTets, minimizeTet(t))
	}

	// use A, B, C, ... etc instead of #
	// max # of tetriminos is 26, so we are safe
	tets := []Tetrimino{}
	letter := 'A'
	for _, t := range cleanedTets {
		tets = append(tets, nameTet(t, letter))
		letter++
	}

	emptyBoard := [][]byte{}
	for i := 0; i < 2; i++ {
		emptyBoard = append(emptyBoard, []byte{})
		for j := 0; j < 2; j++ {
			emptyBoard[i] = append(emptyBoard[i], '.')
		}
	}

	// increase board until it solved
	for {
		if solve(emptyBoard, tets, 0) {
			break
		}
		emptyBoard = boardPlusPlus(emptyBoard)
	}
}

func solve(b [][]byte, tets []Tetrimino, idx int) bool {
	// solved
	if idx >= len(tets) {
		printBoard(b)
		return true
	}

	t := tets[idx]
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			nb, err := putOnBrd(b, t, i, j)
			if err == false {
				if solve(nb, tets, idx+1) {
					return true
				}
			} else {
				// do nothing, tea time ;)
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
				// do nothing, tea time ;)
			} else {
				return nb, true
			}
		}
	}

	return nb, false
}

func parseData(data []byte) ([]Tetrimino, bool) {
	tets := []Tetrimino{}
	var t Tetrimino
	line := ""
	var prev byte
	prev = 46
	for _, b := range data {
		if b == 10 {
			if len(line) != 4 && len(line) != 0 {
				return []Tetrimino{}, false
			}
			if prev == 10 {
				tets = append(tets, t)
				t = Tetrimino{}
			} else {
				// here I cut a bit tet
				// to not include it's empty valid lines
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
	if len(t.data) > 0 {
		tets = append(tets, t)
	}
	return tets, true
}

func PrintTetrimino(tet Tetrimino) {
	for _, s := range tet.data {
		Ft_PutStrLn(s)
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

