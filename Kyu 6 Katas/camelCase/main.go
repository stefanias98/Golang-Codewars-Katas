package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(camelCase("What is love"))
}

func camelCase(s string) string {

	title := strings.Title(s)
	titleWords := strings.Split(title, " ")
	camelCase := strings.Join(titleWords, "")

	return camelCase
}
