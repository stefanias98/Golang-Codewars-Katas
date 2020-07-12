package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ls := "++++++++++++++++++++++++++++++++++++++.++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++++.+++++.++++++++++++.++++++++++++++++++++++.+++++++++++++++++.+++++.++++++.++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++.+.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++.+++++++++++.++++++.++++++++++++++.+++++.+++++++++.+++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++.+++++++++++++.++++++++++++++++++++++++++++++++++++.+++++++.++++++++++++++++++++++.++++++++++++++++.+++++++++++.+.+.++++++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++.+.+++++++++++++++++.++++++++++++++++++++++++++++++++++++.+.++++++++++.++++++++++++++++++++++++++++++++++++++++.+++++++++++++.++++++++++++++.+++++++++++.++++++++++++++++++++++++++++++.++.++++++++++++++++++++++++++++.++++++++++.+++++++++++++++++++++++++"
	//ls := "++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++.+++++++..+++.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++++++++++++++++++++.+++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++."
	//ls := "+++++sdfs++++sdf++++sdklfjsdklfdjmvncxmnxmiuroewurwio+++++++++++++2423423++234+23++234+23++342+2++24++234++++++++++++++???++++++%+++++$#$#++++++++.+++++ssdf+++++++++++++++S+SDFSFSFWWET+BCV+BC+VBN+V+X+++.+++++++.WER.+++.++++++++++++++++++WERW+ERWE++++++++++++XCV+XC++++++++++++++++CXV+XC+XCV++++++++++++++++++++++++XCVXCXCVSTTYJFGDF++++++++++++++++s+dfs+sdf++sdsd+f++SDFS+DFS+FdfsdRTETRCBVCsdfsdf++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++++++++++++++++++++.+++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++."
	fmt.Println(interpreter(ls))
}

func interpreter(s string) (ans string) {

	runes := []rune(s)
	index := 0
	chars := []rune{}
	for _, v := range runes {
		if v == 43 {
			index = (index + 1) % 256
		} else if v == 46 {
			chars = append(chars, rune(index))

		}
	}

	return intToString1(chars)
}

func intToString1(chars []rune) string {

	var ans string
	for _, v := range chars {
		ans += fmt.Sprintf("%02x", v)
	}

	bs, _ := hex.DecodeString(ans)
	return string(bs)
}

func intToString2(chars []rune) string {

	ans := fmt.Sprintf("%02x", chars)
	fmt.Println(ans)
	//ans := "[ 4 4 4]"
	strings.Trim(ans, " ")
	//strings.Trim(strings.Join(strings.Fields(ans), ""), "[]")
	fmt.Println(ans)
	bs, _ := hex.DecodeString(ans)
	return string(bs)
}

func intToString3(chars []int) string {
	b := make([]string, len(chars))
	for i, v := range chars {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, ",")
}