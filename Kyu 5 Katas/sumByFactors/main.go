package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {

	ts := []int{-29804, -4209, -28265, -72769, -31744}
	fmt.Println(sumOfDivided(ts))
}

func sumOfDivided(lst []int) (sol string) {

	primeMap := make(map[int][]int)
	for _, value := range lst {

		check := make(map[int]bool)

		num := 2
		rec := value
		absRec := int(math.Abs(float64(value)))
		for i := 1; i < absRec; i++ {
			if value%num == 0 {
				if check[num] == false {
					primeMap[num] = append(primeMap[num], rec)
					check[num] = true
				}

				value /= num

			} else {
				num++
			}
		}
	}
	fmt.Println(primeMap)
	sortedMap := sorting(primeMap)
	fmt.Println(sortedMap)

	for _, v := range sortedMap {
		sum := 0
		fmt.Println(primeMap[v])
		for _, v := range primeMap[v] {
			sum += v
		}

		sol = sol + fmt.Sprintf("(%s %s)", strconv.Itoa(v), strconv.Itoa(sum))
	}

	return sol
}

func sorting(primeMap map[int][]int) []int {

	keys := make([]int, 0, len(primeMap))
	for k := range primeMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys
}
