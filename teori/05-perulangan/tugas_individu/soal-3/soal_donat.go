package main

import "fmt"

func main() {
	var n, m, x, y int
	fmt.Scan(&n, &m, &x, &y)

	count := 0
	for n >= x && m >= y {
		n -= x
		m -= y
		count++
	}

	fmt.Println(count)
}
