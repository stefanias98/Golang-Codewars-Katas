package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(encode("WEAREDISCOVEREDFLEEATONCE", 3))
	fmt.Println(decode("WECRLTEERDSOEEFEAOCAIVDEN", 3))
}

func encode(s string, n int) string {

	cypherMap := map[int]string{}
	var cypherWord string
	for i, char := range s {
		if (i / (n - 1) % 2) == 0 {
			fmt.Println("Down: ", i%(n-1))
			cypherMap[i%(n-1)] = cypherMap[i%(n-1)] + string(char)
		} else {
			fmt.Println("Up: ", n-1-i%(n-1))
			cypherMap[n-1-i%(n-1)] = cypherMap[n-1-i%(n-1)] + string(char)

		}

	}

	for i := 0; i < n; i++ {
		cypherWord = cypherWord + cypherMap[i]
	}
	return cypherWord
}

func decode(s string, n int) string {

	length := make([]int, len(s))
	sortLength := make([]int, len(s))
	cypherMap := make(map[int]string)
	var decodedString string
	for i := 0; i < len(length); i++ {
		if i%(2*n-2) < n {
			length[i] = i % (2*n - 2)
		} else {
			length[i] = n - 1 - i%(n-1)
		}
	}

	copy(sortLength, length)
	sort.Ints(sortLength)

	for i, value := range sortLength {
		cypherMap[value] += string(s[i])

	}

	for _, value := range length {
		decodedString += string(cypherMap[value][0])
		cypherMap[value] = strings.TrimPrefix(cypherMap[value], string(cypherMap[value][0]))
	}

	return decodedString
}
