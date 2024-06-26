package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//Conditionals

	//If statements

	x := 40
	y := 5

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if y > 10 {
		fmt.Println("Mi alma adora la gloria de papá")
	} else {
		fmt.Println("Esto pasó a otra dimensión")
	}

	if x >= 45 {
		fmt.Println("Wow! que ciencia")
	} else if x == 40 {
		fmt.Println("Heheheheheh")
	} else {
		fmt.Println("Te fundí la masa cefálica")
	}

	//statement; statement

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is greater than or equal x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is less than x which is %v\n", z, x)
	}

	if k := f(); k < x {
		fmt.Printf("k value is %v and that is less than x value = %v\n", k, x)
	} else {
		fmt.Printf("k value is %v and that is greater than x value = %v\n", k, x)
	}
}

func f() int {
	return 2 * rand.Intn(30)
}
