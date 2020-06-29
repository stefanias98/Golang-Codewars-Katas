package main

import (
	"fmt"
	"sort"
	"strings"
)

type ByLen []string

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLen) Less(i, j int) bool { return len(a[i]) > len(a[j]) }

func main() {
	fmt.Println(mix("3qjurRxozuHklfa,khdt2tnth", "Ecqcn>ybtw1poxa?myxqImhwi"))
}

func mix(s1, s2 string) string {

	letterMap1 := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}
	letterMap2 := map[string]int{}
	stringSlice := []string{}

	for k := range letterMap1 {
		letterMap1[k] = strings.Count(s1, k)
		letterMap2[k] = strings.Count(s2, k)

		if letterMap1[k] > 1 || letterMap2[k] > 1 {
			if letterMap2[k] > letterMap1[k] {
				stringSlice = append(stringSlice, "2:"+strings.Repeat(k, letterMap2[k]))
			} else if letterMap1[k] > letterMap2[k] {
				stringSlice = append(stringSlice, "1:"+strings.Repeat(k, letterMap1[k]))
			} else {
				stringSlice = append(stringSlice, "=:"+strings.Repeat(k, letterMap2[k]))
			}
		}

	}

	fmt.Println(stringSlice)
	sort.Sort(ByLen(stringSlice))
	fmt.Println("")
	fmt.Println(stringSlice)

	sameLen := []string{}
	stringSliceSorted := []string{}

	if len(stringSlice) > 1 {
		for i, val := range stringSlice {

			fmt.Println("")
			fmt.Println(len(stringSlice[(i-2+len(stringSlice))%(len(stringSlice)-1)]))
			fmt.Println(len(stringSlice[i]))
			fmt.Println(len(stringSlice[(i+1)%(len(stringSlice))]))
			fmt.Println("")

			if len(stringSlice[i]) == len(stringSlice[(i-2+len(stringSlice))%(len(stringSlice)-1)]) ||
				len(stringSlice[i]) == len(stringSlice[(i+1)%(len(stringSlice))]) {
				sameLen = append(sameLen, val)

				if len(stringSlice[i]) != len(stringSlice[(i+1)%(len(stringSlice))]) || i == len(stringSlice)-1 {
					sort.Strings(sameLen)
					fmt.Println(sameLen)
					stringSliceSorted = append(stringSliceSorted, sameLen...)
					sameLen = []string{}
				}

			} else {
				stringSliceSorted = append(stringSliceSorted, val)

			}

		}
		return strings.Join(stringSliceSorted, "/")
	} else {
		return strings.Join(stringSlice, "/")
	}

}
