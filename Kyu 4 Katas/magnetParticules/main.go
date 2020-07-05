package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(doubles(10, 100))
}

func doubles(maxk, maxn int) float64 {

	var sum float64
	for k := 1; k < maxk+1; k++ {
		fmt.Println("K: ", k)
		for n := 1; n < maxn+1; n++ {
			fmt.Println("N: ", n)

			sum += float64(1) / (float64(k) * math.Pow(float64(n+1), float64(2*k)))
		}
	}

	return sum
}
