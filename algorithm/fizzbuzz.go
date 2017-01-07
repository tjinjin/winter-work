package main

import (
	"flag"
	"fmt"
)

func main() {
	num := flag.Int("num", 0, "num")
	flag.Parse()

	var i int
	for i = 1; i <= *num; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
