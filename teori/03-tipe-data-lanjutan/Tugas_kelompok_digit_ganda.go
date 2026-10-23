package main

import "fmt"

func main() {
	var n int
    fmt.Scan(&n)

    d1 := n / 10
    d2 := n % 10
    hasil := (d1 * 1000) + (d1 * 100) + (d2 * 10) + d2

    fmt.Println(hasil)
}