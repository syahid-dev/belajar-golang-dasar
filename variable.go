package main

import "fmt"

func main() {
	var name string

	name = "Budi"
	fmt.Println(name)

	name = "Badu"
	fmt.Println(name)

	// error :
	// name = 100
	// name = 1.23

	var friendName string = "Tono"
	fmt.Println(friendName)

	var age = 7
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Kafaa Billahi"
		lastName  = "Syahida"
	)
	fmt.Println(firstName, lastName)

	firstName, middleName, lastName := "Kafaa", "Billahi", "Syahida"
	fmt.Println(firstName, middleName, lastName)
}
