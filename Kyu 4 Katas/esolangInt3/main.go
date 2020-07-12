package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(interpreter2("*[s[e]*]", 39, 5, 5))
}

func interpreter2(code string, iterations, width, height int) string {

	dataGrid := make([][]int, height)
	index2 := []int{0}
	skip := []int{0}
	xPos, yPos, validSteps, numOfBracket, catch := 0, 0, 0, 0, 0

	for i := range dataGrid {
		dataGrid[i] = make([]int, width)
	}

	for i := 0; i < len(code); i++ {

		validSteps++
		if validSteps > iterations {
			return getString(dataGrid, height)
		}

		switch {
		case string(code[i]) == "*" && skip[len(skip)-1] == 0:
			dataGrid[yPos][xPos] = (dataGrid[yPos][xPos] + 1) % 2
		case string(code[i]) == "n" && skip[len(skip)-1] == 0:
			yPos = (yPos + height - 1) % height
		case string(code[i]) == "e" && skip[len(skip)-1] == 0:
			xPos = (xPos + width + 1) % width
		case string(code[i]) == "s" && skip[len(skip)-1] == 0:
			yPos = (yPos + height + 1) % height
		case string(code[i]) == "w" && skip[len(skip)-1] == 0:
			xPos = (xPos + width - 1) % width
		case string(code[i]) == "[":
			numOfBracket++
			if skip[len(skip)-1] == 1 {

				validSteps--
				continue

			} else if dataGrid[yPos][xPos] == 0 {

				skip = append(skip, 1)
				catch = numOfBracket

			} else if dataGrid[yPos][xPos] == 1 {

				index2 = append(index2, i)

			}
		case string(code[i]) == "]":
			if skip[len(skip)-1] == 1 && catch == numOfBracket {

				skip = skip[:len(skip)-1]
				validSteps--
				numOfBracket--

			} else if dataGrid[yPos][xPos] == 1 {

				i = index2[numOfBracket]

			} else {

				numOfBracket--
			}

		default:
			validSteps--
		}
	}

	return getString(dataGrid, height)

}

func getString(dataGrid [][]int, height int) (ans string) {
	for i := 0; i < height; i++ {

		ans += strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dataGrid[i])), ""), "[]")

		if i < (height - 1) {

			ans += "\r\n"

		}
	}

	return ans
}
