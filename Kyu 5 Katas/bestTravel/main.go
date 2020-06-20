package main

import (
	"fmt"

	"github.com/gonum/stat/combin"
)

func main() {

	//ts := []int{50, 55, 56, 57, 58}
	//ys := []int{91, 74, 73, 85, 73, 81, 87}
	ks := []int{100, 76, 56, 44, 89, 73, 68, 56, 64, 123, 2333, 144, 50, 132, 123, 34, 89}
	fmt.Println(ks)
	fmt.Println(bestTravel(3760, 40, ks))
}

func bestTravel(t, k int, ls []int) int {

	if (len(ls) < 2) || (k > len(ls)) {
		return -1
	} else {

		comb := combin.Combinations(len(ls), k)
		//fmt.Println(comb)
		combDist := make([][]int, len(comb))
		bestTr := 0
		//var bestDistances []int

		for i := range comb {
			//fmt.Println(comb[i], com)
			for _, val := range comb[i] {
				combDist[i] = append(combDist[i], ls[val])
				//fmt.Println(combDist)
			}
		}

		for i := range combDist {
			travel := 0
			//fmt.Println(combDist[i])
			for j := range combDist[i] {
				travel += combDist[i][j]

			}
			//fmt.Println(travel)
			if travel <= t && travel >= bestTr {
				bestTr = travel
				//bestDistances = combDist[i]
			}
		}
		//fmt.Println(bestDistances)
		if bestTr > 0 {
			return bestTr
		} else {
			return -1
		}

	}
}
