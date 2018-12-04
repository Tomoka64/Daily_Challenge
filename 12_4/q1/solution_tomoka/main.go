package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Println(isCelebratingYear(i))
}

func isCelebratingYear(i int) string {
	if i == 7 || i == 5 || i == 3 {
		return "YES"
	}
	return "NO"
}
