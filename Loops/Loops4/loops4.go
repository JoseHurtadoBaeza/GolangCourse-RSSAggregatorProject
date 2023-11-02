package main

import "fmt"

func fizzbuzz() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 { // Multiples of 3 and 5
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 { // Multiples of 3
			fmt.Println("fizz")
		} else if i%5 == 0 { // Multiples of 5
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

	}

}

// don't touch below this line

func main() {
	fizzbuzz()
}
