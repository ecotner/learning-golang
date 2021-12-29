package main

import (
	"fmt"
	"math"
)

func main() {
	/* primitive data type manipulations */
	fmt.Println("go" + "lang")  // string concatenation with +
	fmt.Println("1 + 1 =", 1+1) // integers and floats
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	/* variables */
	var a string = "initial" // declare variables with `var`
	fmt.Println(a)
	var b, c int = 1, 2 // can declare multiple variables at once
	fmt.Println(b, c)
	var d = true // go will infer the data type on assignment (bool in this case)
	fmt.Println(d)
	var e int // when variable is declared but not assigned, it is "zero-valued" (in this case literally zero)
	fmt.Println(e)
	f := "apple" // `f :=` is shorthand for `var f =`
	fmt.Println(f)

	/* constants */
	{ // you can scope variables using curly braces - notice how `d` is defined twice but it doesn't cause a conflict
		const s string = "constant" // `const` declares a constant primitive type
		const n = 500000000         // they can perform arithmetic with arbitrary precision
		const d = 3e20 / n
		fmt.Println(d)
		fmt.Println(int64(d))    // a numeric constant has no type until it is given one, such as by an explicit conversion
		fmt.Println(math.Sin(n)) // or by using it in a context that requires a conversion
	}
	fmt.Println(d) // this is just to show that `d` still has its original value outside the scoped block above
}
