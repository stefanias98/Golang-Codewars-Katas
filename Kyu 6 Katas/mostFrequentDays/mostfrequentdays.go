package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(test(2427))
	fmt.Println(test(2185))
	fmt.Println(test(1167))
	fmt.Println(test(1770))
	fmt.Println(test(1785))
	fmt.Println(test(1984))
	fmt.Println(test(3076))

	//  test(2427, "Friday")
	//	test(2185, "Saturday")
	//	test(1167, "Sunday")
	//	test(1770, "Monday")
	//	test(1785, "Saturday")
	//	test(1984, "Monday", "Sunday")
	//	test(3076, "Saturday", "Sunday")
}

func test(year int) []string {
	var mfd []string
	//yearF := math.Mod(float64(year)/100, 4)
	//fmt.Println("")
	//fmt.Println(year)
	//fmt.Println(year%4 != 0)
	//fmt.Println((year%4 == 0 && yearF == math.Trunc(yearF)))
	//fmt.Println(yearF)
	//fmt.Println(math.Trunc(yearF))

	if year%400 == 0 || year%4 == 0 && year%100 != 0 {

		t1 := time.Date(year, time.December, 30, 23, 59, 59, 0, time.UTC)
		t2 := time.Date(year, time.December, 31, 23, 59, 59, 0, time.UTC)

		dayIndex1 := int(t1.Weekday())
		dayIndex2 := int(t2.Weekday())

		//fmt.Println(dayIndex1, dayIndex2)

		dayIndex1 = (dayIndex1 + 6) % 7
		dayIndex2 = (dayIndex2 + 6) % 7

		//fmt.Println(dayIndex1, dayIndex2)

		if dayIndex1 < dayIndex2 {
			mfd = append(mfd, t1.Weekday().String(), t2.Weekday().String())
		} else {
			mfd = append(mfd, t2.Weekday().String(), t1.Weekday().String())
		}

		return mfd

	} else {

		t1 := time.Date(year, time.December, 31, 23, 59, 59, 0, time.UTC)
		mfd = append(mfd, t1.Weekday().String())
		return mfd
	}
}
