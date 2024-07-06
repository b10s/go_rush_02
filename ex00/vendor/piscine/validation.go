package piscine

import (
	"os"
	"fmt"
// "ft"
)

func Ft_ValidateInput(fileName string) bool {
	return true
}

func Ft_ValidateInput1(fileName string) bool {
	byte_slice, err := os.ReadFile(fileName)
	if err != nil {
		Ft_PutStrLn("error: read\n")
		//	remove fmt
		fmt.Println(err)
//		piscine.Ft_PutStrLn("error: read\n")
//		piscine.Ft_PutStr(err)
		return false
	}

	Ft_PutStr("----@Ft_ValidateInput----\n")
	Ft_PutStr(string(byte_slice))
	Ft_PutStr("-------------------------\n")
//	piscine.Ft_PutStr("----@Ft_ValidateInput----\n")
//	piscine.Ft_PutStr(s)
//	piscine.Ft_PutStr("-------------------------\n")

// do we have to cast string to rune slice ? We have to take care of
// in case of multi-byte characters are given as parameter.

//	data := []rune(byte_slice)
	data := []rune(string(byte_slice))

//	for _, c := range s {
//		data :=
//	}

	if IsValidElement(data) {
		Ft_PutStr("error: invalid element.\n")
//		piscine.Ft_PutStr("error: invalid element.\n")
		return false
	}



	if IsValidSize(data) {
		Ft_PutStr("error: invalid size.\n")
//		piscine.Ft_PutStr("error: invalid size.\n")
		return false
	}



	if IsBlockCharAdjacentAnotherBlockChar(data) {
		Ft_PutStr("error: invalid size.\n")
//		piscine.Ft_PutStr("error: invalid size.\n")
		return false
	}


	return true
}



func IsValidElement(s []rune) bool {
	for _, c := range s {
		if c != '#' && c != '.' {
			return false
		}
	}
	return true
}

// func isValidChar(c rune) bool {
// 	if c == '#' || c == '.' {
// 		return true
// 	}
// 	return false
// }

// func isBlockChar(c rune) bool {
// 	if c == '#' {
// 		return true
// 	}
// 	return false
// }

// func isBlankChar(c rune) bool {
// 	if c == '.' {
// 		return true
// 	}
// 	return false
// }



func skip4x4map(data []rune) bool {
	for height := 0; height <= 3; height++ {
		if data[height] == '\n' {
		}
		for width := 0; width <= 3; width++ {
			return true
		}
	}
	return true
}

func IsValidSize(data []rune) bool {
	if isValidHeight(data) && isValidWidth(data) {
		return true
	}
	return false
}

// if 4, true
func isValidHeight(data []rune) bool {
//	for _, c := range data {
//		if
//		return false
//	}
	return false
}

// if 4, true
func isValidWidth(data []rune) bool {
	return true
}

// if 1, true
func isSeparatedByOneNewLine(data []rune) bool {
	return true
}



// if 4, true
func IsValidNumBlockChar() bool {
	return true
}



// dfs
func IsBlockCharAdjacentAnotherBlockChar(data []rune) bool {
	return true
}
