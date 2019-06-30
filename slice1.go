package main

import (
	"fmt"
	"strings"
)

func main() {
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	game[0][0] = "X"
	game[1][1] = "X"
	game[2][2] = "X"
	game[2][0] = "O"
	game[0][2] = "O"
	game[3][2] = "O"

	fmt.Println(game)
	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " | "))
	}
}
