package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sayHello(sentence string, z int) {

	fmt.Println(sentence, z)

}

func main() {
	fmt.Println(add(5, 6))
	sayHello("Kelvin was here", 5)
}
