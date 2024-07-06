package main

// remove fmt from here
import (
	"os"
	"piscine"
)

func main() {
	// TODO: check if number of args is correct
	//	if not, print usage
	if len(os.Args) != 2 {
		piscine.Ft_PutUsage()
		return
	}

	// TODO: validate
	fileName := os.Args[1]
	piscine.Ft_PutStrLn(fileName)
	if !piscine.Ft_ValidateInput(fileName) {
		piscine.Ft_PutStrLn("error")
		return
	}

	piscine.Ft_solve(fileName)
}
