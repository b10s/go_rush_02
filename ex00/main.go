package main

// remove fmt from here
import (
	"fmt"
	"os"
)

// TODO: move to utils package?
// 	write there something like print string
// all start with ft
func ft_PrintUsage() {
	fmt.Println("usage")
}

func ft_ValidateInput(fileName string) bool {
	return true
}

func main() {
	// TODO: check if number of args is correct
	//	if not, print usage
	if (len(os.Args) !=2 ) {
		ft_PrintUsage()
		return
	}

	// TODO: validate
	fileName := os.Args[1]
	fmt.Println(fileName)
	if !ft_ValidateInput(fileName) {
		fmt.Println("Error")
		return
	}

	ft_solve(fileName)
}

func ft_solve(fileName string) {
}
