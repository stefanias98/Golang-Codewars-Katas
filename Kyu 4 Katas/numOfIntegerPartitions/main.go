package main

import "fmt"

func main() {
	fmt.Println(partitions(50))
}

func partitions(n int) int {

	ans := make([]int, n+1)
	ans[0] = 1

	for i := 0; i < n; i++ {
		temp := make([]int, n-i)
		for j := range temp {
			ans[n-len(temp)+j+1] += ans[j]
		}
	}
	return ans[n]
}
