package main

import "fmt"

func main() {
	fmt.Println(josephusSurvivor(7, 3))
}

func josephusSurvivor(n, k int) int {
	result := []int{}
	inter := []int{}

	for i := 0; i < n; i++ {
		inter = append(inter, i+1)
	}

	length := len(inter)
	var kt int
	for i := 1; i <= length; i++ {
		kt = ((kt - 1 + k) % len(inter))
		result = append(result, inter[kt])
		inter = append(inter[:kt], inter[kt+1:]...)
	}
	fmt.Println(result[len(result)-1])

	return result[len(result)-1]
}
