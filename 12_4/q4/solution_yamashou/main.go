package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	allFactors := make(map[int]int)

	for i := 2; i <= n; i++ {
		factors := pf(i)

		for k, v := range factors {
			_, ok := allFactors[k]
			if ok {
				allFactors[k] += v
			} else {
				allFactors[k] = v
			}
		}
	}

	c2, c4, c14, c24, c74 := countNum(allFactors)

	ans := c4*
		(c4-1)*
		(c2-2)/2 +
		c24*(c2-1) +
		c14*(c4-1) +
		c74

	fmt.Println(ans)
}

func countNum(factors map[int]int) (c2, c4, c14, c24, c74 int) {
	for _, v := range factors {
		if v >= 2 {
			c2++
		}
		if v >= 4 {
			c4++
		}
		if v >= 14 {
			c14++
		}
		if v >= 24 {
			c24++
		}
		if v >= 74 {
			c74++
		}
	}

	return
}

func pf(num int) map[int]int {
	if num == 1 {
		return make(map[int]int)
	}

	ret := make(map[int]int)
	maxPrime := num
	for num%2 == 0 {
		ret[2]++
		num /= 2
	}

	for i := 3; i*i <= maxPrime; i = i + 2 {
		if num == 1 {
			break
		}

		for num%i == 0 {
			_, ok := ret[i]
			if ok {
				ret[i]++
			} else {
				ret[i] = 1
			}
			num /= i
		}
	}

	if num > 2 {
		_, ok := ret[num]
		if ok {
			ret[num]++
		} else {
			ret[num] = 1
		}
	}

	return ret
}
