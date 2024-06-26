package main

import "fmt"

//Function that run first that any others intruction using Go
func init() {
	fmt.Println("This will run first than any code on main func()")
}

func main() {
	//SEQUENCE
	fmt.Println("First statement to run")
	fmt.Println("Second statement to run")
	x := 40 //Third Statement to run
	y := 5  //Fourth Statement to run
	fmt.Printf("x=%v\ny=%v", x, y)
}
