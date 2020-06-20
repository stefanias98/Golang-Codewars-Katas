package main

import "fmt"

func main() {
	fmt.Println(mixbonacci([]string{"pel"}, 10))
}

func mixbonacci(pattern []string, length int) []int64 {

	if len(pattern) == 0 {
		return []int64{}
	} else {
		seq := make([]int64, 0)
		patRecord := map[string]int{
			"fib": 0,
			"pad": 0,
			"jac": 0,
			"pel": 0,
			"tri": 0,
			"tet": 0,
		}

		funcRecord := map[string]func(int) int64{
			"fib": fibonacci,
			"pad": padovan,
			"jac": jacobsthal,
			"pel": pell,
			"tri": tribonacci,
			"tet": tetranacci,
		}

		for i := 0; i < length; i++ {
			curPat := pattern[i%len(pattern)]
			seq = append(seq, funcRecord[curPat](patRecord[curPat]))
			patRecord[curPat]++
		}

		return seq
	}
}

func fibonacci(n int) int64 {
	n2, n1, ncurr := 0, 0, 1

	if n <= 1 {
		return int64(n)
	} else {
		for i := 1; i < n; i++ {
			n2 = n1
			n1 = ncurr
			ncurr = n1 + n2
		}
		return int64(ncurr)
	}
}

func padovan(n int) int64 {

	n3, n2, n1, ncurr := 1, 0, 0, 1

	if n == 0 {
		return int64(1)
	} else if n == 1 || n == 2 {
		return int64(0)
	} else {
		for i := 3; i < n; i++ {
			n3 = n2
			n2 = n1
			n1 = ncurr
			ncurr = n2 + n3
		}
		return int64(ncurr)
	}
}

func jacobsthal(n int) int64 {
	n2, n1, ncurr := 0, 0, 1

	if n <= 1 {
		return int64(n)
	} else {
		for i := 1; i < n; i++ {
			n2 = n1
			n1 = ncurr
			ncurr = n1 + 2*n2
		}
		return int64(ncurr)
	}
}

func pell(n int) int64 {
	n2, n1, ncurr := 0, 0, 1

	if n <= 1 {
		return int64(n)
	} else {
		for i := 1; i < n; i++ {
			n2 = n1
			n1 = ncurr
			ncurr = 2*n1 + n2
		}
		return int64(ncurr)
	}
}

func tribonacci(n int) int64 {

	n3, n2, n1, ncurr := 0, 0, 1, 1

	if n < 2 {
		return 0
	} else if 2 <= n && n < 4 {
		return int64(1)
	} else {
		for i := 3; i < n; i++ {
			n3 = n2
			n2 = n1
			n1 = ncurr
			ncurr = n3 + n2 + n1
		}

		return int64(ncurr)
	}

}

func tetranacci(n int) int64 {

	n4, n3, n2, n1, ncurr := 0, 0, 0, 1, 1

	if n < 3 {
		return 0
	} else if 3 <= n && n < 5 {
		return int64(1)
	} else {
		for i := 4; i < n; i++ {
			n4 = n3
			n3 = n2
			n2 = n1
			n1 = ncurr
			ncurr = n4 + n3 + n2 + n1
		}

		return int64(ncurr)
	}
}
