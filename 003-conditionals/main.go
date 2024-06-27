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

	//Switch statements

	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL TO 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("default value here")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	case 44:
		fmt.Println("x is 44")
	default:
		fmt.Println("default value here 2")
	}

	// no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}

	// no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 43")
		fallthrough
	default:
		fmt.Println("printed because of ALL OF THE fallthrough statements: this is the default case for x")
	}

}

func f() int {
	return 2 * rand.Intn(30)
}
