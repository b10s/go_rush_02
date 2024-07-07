package piscine

//TODO: remove fmt
import (
	"os"
	"fmt"
)

func Ft_ValidateInput(fileName string) bool {
	fmt.Println("validation ...")

	data, err := os.ReadFile(fileName)
	if err != nil {
		return false
	}

	// no empty files
	// 5 is to be safe, i need to look back few bytes
	if len(data) < 5 {
		return false
	}

	// check if there is invalid char
	// 10 - end of line
	// 35 - #
	// 46 - .
	for _, b := range data {
		if b != 46 && b != 35 && b!= 10 {
			return false
		}
	}

	// no files starting from new line
	if data[0] == 10 {
		return false
	}

	// no two empty lines in a row
	nlcnt := 0
	for _, b := range data {
		if b == 10 {
			nlcnt++
		} else {
			nlcnt = 0
		}
		if nlcnt > 2 {
			return false
		}
	}

	// no new line at the end
	if data[len(data) - 1] == 10 && data[len(data) - 2] == 10 {
		return false
	}


	// check shape is okay
	// line len == 4
	line_len := 0
	for _, b := range data {
		if b == 10 {
			if line_len != 4 && line_len != 0 {
				return false
			}
			line_len = 0
		} else {
			line_len++
		}
	}

	// tet height is 4 lines
	pre := byte(42)
	line_cnt := 0
	for _, b := range data {
		// no huge tetros
		if line_cnt > 4 {
			return false
		}
		// it is empty line
		if b == 10 {
			if pre == 10 {
				if line_cnt != 4 {
					return false
				}
				line_cnt = 0
			} else {
				line_cnt++
			}
		}
		pre = b
	}

	// make sure all tets are valid tets
	// sum of each tet neighbors should be 6: no less, no more
	tet := [][]byte{}
	line := []byte{}
	pre = 42
	for _, b := range data {
		if b == 10 {
			// empty line, tet completed
			// check tet
			if pre == 10 {
				if countBlocks(tet) != 4 {
					return false
				}
				if countAdjBlcks(tet) < 6 {
					return false
				}
				tet = [][]byte{}
				line = []byte{}
			// just another line of tet
			// keep building
			} else {
				tet = append(tet, line)
				line = []byte{}
			}
		} else {
			line = append(line, b)
		}
		pre = b
	}

	//check last tet
	if countBlocks(tet) != 4 {
		return false
	}
	if countAdjBlcks(tet) < 6 {
		return false
	}


	return true
}


func countBlocks(tet [][]byte) int {
	res := 0
	for i:=0; i<len(tet); i++ {
		for j:=0; j<len(tet[i]); j++ {
			if tet[i][j] == 35 {
				res++
			}
		}
	}
	return res
}

func countAdjBlcks(tet [][]byte) int {
	res := 0
	for i:=0; i<len(tet); i++ {
		for j:=0; j<len(tet[i]); j++ {
			if tet[i][j] == 35 {
				// upper block
				if i - 1 >=0 {
					if tet[i-1][j] == 35 {
						res++
					}
				}
				// lower block
				if i + 1 < len(tet) {
					if tet[i+1][j] == 35 {
						res++
					}
				}
				// left block
				if j - 1 >= 0 {
					if tet[i][j-1] == 35 {
						res++
					}
				}
				// right block
				if j+ 1 < len(tet[i]) {
					if tet[i][j+1] == 35 {
						res++
					}
				}
			}
		}
	}
	return res
}
