package main

import (
	"fmt"
	"time"
)

func main() {
	/* if/else/else if */
	if 7%2 == 0 { // basic if/else branch
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // an `if` without an `else`
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // `num` is scoped to this conditional block
		fmt.Println(num, "is negative")
	} else if num < 10 { // can use `else if` to test additional logic
		fmt.Println("has 1 digit")
	} else {
		fmt.Println("has multiple digits")
	}

	/* for loops */
	i := 1
	for i <= 3 { // single loop condition; equivalent to a while loop in other languages
		fmt.Println(i)
		i += 1 // iteration occurs inside the loop
	}

	for j := 7; j <= 9; j++ { // initialization, loop condition, and iteration all happens in one line
		fmt.Println(j)
	}

	for { // naively loops forever
		fmt.Println("loop")
		break // breaks the loop
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 { // only triggers if `n` is even
			continue // skips the rest of the loop body and goes to the next iteration
		}
		fmt.Println(n)
	}

	/* switch/case */
	ii := 2
	fmt.Print("Write ", ii, " as ")
	switch ii {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // you can use commas to test multiple things
		fmt.Println("It's the weekend")
	default: // use `default` if it doesn't match any previous cases
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch { // this is basically an alternate way to express if/else
	case t.Hour() < 12: // case can be any expression that evaluates to a bool
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) { // function for checking types with a switch statement
		switch t := i.(type) { // `t` is scoped to the switch block (not that it matters since it's also scoped to the function)
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
