package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(interpreter("*>*>>>*>*>>>>>*[>*]", "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"))
}

func interpreter(code, tape string) string {

	index := 0
	record := make([]int, len(tape))
	for _, v := range tape {
		switch string(v) {
		case "0":
			record[index] = 0
			index++
		case "1":
			record[index] = 1
			index++
		}
	}
	fmt.Println(record)
	index = 0
	index2 := []int{0}
	numofBracket := 0
	skip := []int{0}
	catch := 0
	fmt.Println(skip)
	for i := 0; i < len(code); i++ {
		fmt.Println(i)
		if index > (len(tape)-1) || index < 0 {
			return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(record)), ""), "[]")
		}

		switch {
		case string(code[i]) == "*" && skip[len(skip)-1] == 0:
			record[index] = (record[index] + 1) % 2
		case string(code[i]) == ">" && skip[len(skip)-1] == 0:
			index++
		case string(code[i]) == "<" && skip[len(skip)-1] == 0:
			index--
		case string(code[i]) == "[":
			numofBracket++
			if skip[len(skip)-1] == 1 {
				continue
			} else if record[index] == 0 {

				skip = append(skip, 1)
				catch = numofBracket

			} else if record[index] == 1 {

				index2 = append(index2, i)
				fmt.Println(index2)
				fmt.Println("Breee2", numofBracket, index2[numofBracket])
			}
		case string(code[i]) == "]":
			//fmt.Println(len(skip) - 1)
			if skip[len(skip)-1] == 1 && catch == numofBracket {
				skip = skip[:len(skip)-1]

			} else if record[index] == 1 {
				fmt.Println("Breeeeee", numofBracket, index2[numofBracket])
				i = index2[numofBracket] - 1
				fmt.Println(i)
				//numofBracket--
			}
			numofBracket--
		}
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(record)), ""), "[]")
}

/*
0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
1100110000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000

*/
