package main

import (
	"fmt"
	"strconv"
)

func main() {

	ls := []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}
	fmt.Println(solution(ls))
}

func solution(list []int) string {
	ans := ""
	store1 := -9999
	store2 := -9999
	list = append(list, 0)
	list = append(list, 0)
	for i := range list {
		if (i < len(list)-2) && list[i+2]-list[i] == 2 && store1 == -9999 {
			store1 = list[i]
		} else if (i < len(list)-1) && list[i+1]-list[i] != 1 && store1 != -9999 {
			store2 = list[i]
			ans += strconv.Itoa(store1) + "-" + strconv.Itoa(store2) + ","
			store1, store2 = -9999, -9999
		} else if (i < len(list)-2) && list[i+2]-list[i] != 2 && store1 == -9999 {
			ans += strconv.Itoa(list[i]) + ","
		}
	}
	return ans[:len(ans)-1]
}
