package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(rectangleRotation(30, 2))
}

func rectangleRotation(a, b int) int {

	div := math.Sqrt(2) / 2

	numA := (float64(a) / 2) / div
	numB := (float64(b) / 2) / div

	intNumA, intNumB := 2*int(numA)+1, 2*int(numB)+1

	aa := (intNumA - intNumA/2)
	bb := (intNumB / 2)
	cc := (intNumB - intNumB/2)
	dd := (intNumA / 2)

	if intNumA%4 == 1 && intNumB%4 == 3 {
		return aa*bb + cc*dd
	} else if intNumA%4 == 3 && intNumB%4 == 1 {
		return aa*bb + cc*dd
	} else {
		return aa*cc + bb*dd
	}
}
