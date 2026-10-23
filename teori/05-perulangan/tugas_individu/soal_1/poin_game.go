package main

import "fmt"

func main() {
	var poin, total int

	for {
		fmt.Scan(&poin)
		if poin == 0 {
			break
		}
		total += poin
	}

	fmt.Println(total)
}
