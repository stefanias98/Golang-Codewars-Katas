package main

import (
	"fmt"
	"strings"
)

func main() {

	//Sample tests
	fmt.Println(validBraces("{}[]()][]()")) //false
	fmt.Println(validBraces("{}]"))         //false
	fmt.Println(validBraces("(){}[]"))      //true
	fmt.Println(validBraces("([{}])"))      //true
	fmt.Println(validBraces("(}"))          //false
	fmt.Println(validBraces("[(])"))        //false
	fmt.Println(validBraces("[({)](]"))     //false
	fmt.Println(validBraces("(({[]}))"))    //true
	fmt.Println(validBraces("{{{{]}}}}"))   //false

}

func validBraces(str string) bool {

	split := strings.Split(str, "")

	inputBrackets := []string{}
	openBrackets := []string{"(", "[", "{"}
	closedBrackets := []string{")", "]", "}"}

	for i, char := range split {

		if char == openBrackets[0] || char == openBrackets[1] || char == openBrackets[2] {
			inputBrackets = append(inputBrackets, char)

		} else if char == closedBrackets[0] || char == closedBrackets[1] || char == closedBrackets[2] {
			if i == 0 {
				return false
			} else {
				if len(inputBrackets) == 0 {
					return false
				} else if len(inputBrackets) > 0 {
					index := len(inputBrackets) - 1

					if indexOf(char, closedBrackets) != indexOf(inputBrackets[index], openBrackets) {
						return false
					} else {

						inputBrackets = inputBrackets[:len(inputBrackets)-1]

					}
				} else {
					continue
				}

			}

		}

	}
	if len(inputBrackets) == 0 {
		return true
	} else {
		return false
	}

}

func indexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
