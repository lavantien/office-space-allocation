package main

import (
	"fmt"
	"io"
	"os"
)

const (
	totalLevel    = 4
	unitsPerLevel = 12
)

func main() {
	spaces := ReadSpaceMap(os.Stdin, totalLevel, unitsPerLevel)
	// debugPrintSpaces(spaces)

	result := Allocator(totalLevel, unitsPerLevel, spaces)

	fmt.Println(result)
}

func Allocator(totalLevel int, unitsPerLevel int, spaces [][]bool) []string {
	var res []string

	return res
}

func ReadSpaceMap(inSource io.Reader, totalLevel int, unitsPerLevel int) [][]bool {
	spaces := make([][]bool, 0)

	for i := 0; i < totalLevel; i++ {
		spaces = append(spaces, make([]bool, 0))
		for j := 0; j < unitsPerLevel; j++ {
			var t bool
			fmt.Fscan(inSource, &t)
			spaces[i] = append(spaces[i], t)
		}
	}

	return spaces
}

func debugPrintSpaces(spaces [][]bool) {
	for i := 0; i < len(spaces); i++ {
		for j := 0; j < len(spaces[i]); j++ {
			fmt.Print(spaces[i][j], " ")
		}
		fmt.Println()
	}
}
