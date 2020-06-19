package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findTheSmallest(80424994128272))
}

func findTheSmallest(n int64) []int64 {

	ans := make([]int64, 3)
	if n < 10 {
		ans[0] = n
		return ans
	}

	numberString := strconv.Itoa(int(n))
	digitString := strings.Split(numberString, "")
	fmt.Println(digitString)

	var digits []int
	for _, digit := range digitString {
		num, _ := strconv.Atoi(digit)
		digits = append(digits, num)
	}

	sortedDig := make([]int, len(digits))
	ansDig := make([]int, 0)
	copy(sortedDig, digits)
	sort.Ints(sortedDig)
	var minAvDig int
	index := 0
	var found bool

	fmt.Println(digits)
	fmt.Println(sortedDig)
	for i := range digits {
		if digits[i] != sortedDig[i] && found == false {
			minAvDig = sortedDig[i]
			found = true
		}

		if digits[i] == minAvDig {
			ans[1] = int64(i)
		}
	}

	found = false
	fmt.Println(digits)
	for i := range digits {

		if digits[i] >= minAvDig {
			index = i
			break
		}

	}

	fmt.Println(digits[ans[1]])

	fmt.Println("Dig :", digits[index:ans[1]])
	ansDig = append(ansDig, digits[:index]...)

	fmt.Println(ansDig)
	fmt.Println("Dig :", digits[index:ans[1]])
	ansDig = append(ansDig, minAvDig)
	fmt.Println(ansDig)
	fmt.Println("Dig :", digits[index:ans[1]])
	ansDig = append(ansDig, digits[index:ans[1]]...)
	fmt.Println(ansDig)
	fmt.Println("Dig :", digits[index:ans[1]])
	ansDig = append(ansDig, digits[ans[1]+1:]...)
	fmt.Println(ansDig)

	ans[2] = int64(index)

	if (ans[1]-ans[2]) == -1 || (ans[1]-ans[2]) == 1 {
		ans[1], ans[2] = ans[2], ans[1]
	}

	ansStr := strings.Trim(strings.Replace(fmt.Sprint(ansDig), " ", "", -1), "[]")
	fmt.Println(ansStr)

	if strings.HasPrefix(ansStr, "0") {
		ansNum, _ := strconv.Atoi(ansStr[1:])
		ans[0] = int64(ansNum)
		return ans
	} else {
		ansNum, _ := strconv.Atoi(ansStr)
		ans[0] = int64(ansNum)
		return ans
	}

}

/*
275468899
275 4 68899
2 4 75 68 899
247568899

261235
26 1 235
126 235
126235

209917
029917

187863002809


*/
