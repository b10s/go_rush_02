package main

import (
	"os"
	"piscine"
)

func main() {
	if len(os.Args) != 2 {
		piscine.Ft_PutUsage()
		return
	}

	fileName := os.Args[1]
	if !piscine.Ft_ValidateInput(fileName) {
		piscine.Ft_PutStrLn("error")
		return
	}

	piscine.Ft_solve(fileName)
}
