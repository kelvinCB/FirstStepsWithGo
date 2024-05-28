package main

import "fmt"

// single type function
func add(x, y int) int {
	return x + y
}

// multiple type function
func sayHello(sentence string, z int) {

	fmt.Println(sentence, z)

}

// multiple results function
func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	fmt.Println(add(5, 6))
	sayHello("Kelvin was here", 5)

	a, b := swap("Hello", "Go")
	fmt.Println(a, b)
}
