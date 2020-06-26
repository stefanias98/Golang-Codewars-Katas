package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(nextBigger(435894875))
}

func nextBigger(n int) int {

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

		if i == 0 {
			return -1
		} else if digits[i] > digits[i-1] {
			deltaNum := digits[i:]
			sort.Ints(deltaNum)

			for j, num := range deltaNum {
				if num > digits[i-1] {
					digits[i-1], deltaNum[j] = deltaNum[j], digits[i-1]
					digits = append(digits[:i], deltaNum...)
					break
				}
			}

			break
		}

	}

	ansNum, _ := strconv.Atoi(strings.Trim(strings.Replace(fmt.Sprint(digits), " ", "", -1), "[]"))

	return ansNum
}

/*495
4 95
4 59
5 49


4951
4 951
4 159
5 149

5149

435894875
435894 875
435894 578
435895 478
435895478

583058309523505483250
5830583095235054832 50
5830583095235054832 05
5830583095235054835 02
583058309523505483502*/
