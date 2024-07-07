package piscine

//TODO: remove fmt
import (
	"os"
	"fmt"
)

func Ft_ValidateInput(fileName string) bool {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return false
	}

	// check if there is invalid char
	fmt.Println(data)
	for _, b := range data {
		if b != 46 && b != 35 && b!= 10 {
			return false
		}
	}

// make check for new lines

// make chekc shape is okay

// make sure all tets are valid tets

	return true
}


