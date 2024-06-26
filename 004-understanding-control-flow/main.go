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

	// Operators

	//Comparison Operators

	/*
	   == equal
	   != not equal
	   < less
	   <= less or equal
	   > greater than
	   >= greater or equal than

	*/

	//Arithmetic operators

	/*
		+    sum                    integers, floats, complex values, strings
		-    difference             integers, floats, complex values
		*    product                integers, floats, complex values
		/    quotient               integers, floats, complex values
		%    remainder              integers

		&    bitwise AND            integers
		|    bitwise OR             integers
		^    bitwise XOR            integers
		&^   bit clear (AND NOT)    integers

		<<   left shift             integer << integer >= 0
		>>   right shift            integer >> integer >= 0
	*/

	//Logical Operators

	/*
		&&    conditional AND    p && q  is  "if p then q else false"
		||    conditional OR     p || q  is  "if p then true else q"
		!     NOT                !p      is  "not p"
	*/

	//Address operators

	/*
		&x
		&a[f(2)]
		&Point{2, 3}
		*p
		*pf(x)

		var x *int = nil
		*x   // causes a run-time panic
		&*x  // causes a run-time panic

	*/

}
