package main

import "fmt"

func main() {

	a := []int{}
	inter := []interface{}{}

	for i := range a {
		inter = append(inter, i+1)
	}
	fmt.Println(inter)
	fmt.Println("")

	fmt.Println(josephus(inter, 3))
}

func josephus(items []interface{}, k int) []interface{} {
	result := make([]interface{}, 0)
	fmt.Println("RES: ", result)
	if &items == &result {
		return result
	} else {
		length := len(items)
		var kt int
		for i := 1; i <= length; i++ {
			kt = ((kt - 1 + k) % len(items))
			result = append(result, items[kt])
			items = append(items[:kt], items[kt+1:]...)
		}
		return result
	}
}
