package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {

	num := flag.Int("num", 0, "num")
	flag.Parse()

	var i int
	for i = 1; i <= *num; i++ {
		var s string
		if i%3 == 0 {
			s = "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		fmt.Println(s)
	}
}
