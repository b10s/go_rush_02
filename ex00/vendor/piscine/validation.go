package piscine

import (
	"os"
)

func Ft_ValidateInput(fileName string) bool {
	_, err := os.ReadFile(fileName)
	if err != nil {
		return false
	}


	return true
}


// make check for new lines

// make check only valid chars are here

// make chekc shape is okay

// make sure all tets are valid tets

// make sure thera no more than 26 tets (maybe do it in sove since I can parse all tets there and easily fail there)

