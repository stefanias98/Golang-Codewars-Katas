package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(nextSmaller(1027))
}

func nextSmaller(n int) int {

	if n < 10 {
		return -1
	}

	numberString := strconv.Itoa(n)
	digitString := strings.Split(numberString, "")
	fmt.Println(digitString)

	var digits []int
	for _, digit := range digitString {
		num, _ := strconv.Atoi(digit)
		digits = append(digits, num)
	}
	fmt.Println(digits)

	for i := len(digits) - 1; i >= 0; i-- {

		fmt.Println(digits[i])
		if i == 0 {
			return -1
		} else if digits[i] < digits[i-1] {
			deltaNum := digits[i:]
			sort.Ints(deltaNum)

			for j, num := range deltaNum {
				if num < digits[i-1] {
					digits[i-1], deltaNum[j] = deltaNum[j], digits[i-1]
					sort.Sort(sort.IntSlice(deltaNum))
					digits = append(digits[:i], deltaNum...)
					break
				}
			}

			break
		}

	}

	ansStr := strings.Trim(strings.Replace(fmt.Sprint(digits), " ", "", -1), "[]")

	if strings.HasPrefix(ansStr, "0") {
		return -1
	} else {
		ansNum, _ := strconv.Atoi(ansStr)
		return ansNum
	}

}

/*
495

49 5
45 9
459

4951
495 1
491 5
4915

21345
2 1345
1 2345
1 5432
15432

1027
1 027
0 127
0 721
0721
*/
