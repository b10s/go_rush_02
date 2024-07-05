package main

// remove fmt from here
import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	// TODO: check if number of args is correct
	//	if not, print usage
	if len(os.Args) != 2 {
		piscine.Ft_PrintUsage()
		return
	}

	// TODO: validate
	fileName := os.Args[1]
	fmt.Println(fileName)
	if !piscine.Ft_ValidateInput(fileName) {
		fmt.Println("Error")
		return
	}

	ft_solve(fileName)
}

func ft_solve(fileName string) {
}
