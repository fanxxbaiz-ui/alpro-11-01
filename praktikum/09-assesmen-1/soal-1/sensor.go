package main

import (
	"fmt"
)

func main() {
	var data []int
	var x int

	for fmt.Scan(&x); ; fmt.Scan(&x) {
		data = append(data, x)
		n := len(data)
		if n >= 2 && data[n-1] == data[n-2] {
			break
		}
	}

	data = data[:len(data)-1]
	if len(data) == 0 {
		return
	}

	max, min, sum := data[0], data[0], 0
	for _, v := range data {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
		sum += v
	}

	avg := float64(sum) / float64(len(data))
	fmt.Printf("%d %d %.2f\n", max, min, avg)
}
