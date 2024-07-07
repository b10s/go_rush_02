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

	return true
}


