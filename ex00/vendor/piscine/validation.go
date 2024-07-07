package piscine

import (
	"os"
)

func Ft_ValidateInput(fileName string) bool {
	_, err := os.ReadFile(fileName)
	if err != nil {
		Ft_PutStrLn("error: read\n")
		return false
	}

	return true
}
