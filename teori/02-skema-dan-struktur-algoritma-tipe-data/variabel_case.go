package main

import "fmt"

func main() {
	var name string

	name = "Fathan Al Akbar"
	fmt.Println("nama : ", name)

	var lastName = "Akbar"
	fmt.Println("Nama Belakang : ", lastName)

	middleName := "Al"
	fmt.Println("Nama Tengah: ", middleName)

	var (
		fullName = "Fathan Al Akbar"
		firstName = "Fathan"
	)

	fmt.Println(fullName)
	fmt.Println(firstName)
}