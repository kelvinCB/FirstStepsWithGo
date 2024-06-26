package main

import "fmt"

//Default empty variables values

var c, python, java bool
var i, j int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)
	initializers()
	short()
}

// Variables with initializers

func initializers() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

// Short variable declarations

func short() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
