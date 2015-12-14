package main

import (
	"fmt"
	bnlmz "github.com/coffeeSouffle/bnlmz"
	"os"
)

func main() {
	num := os.Args
	fmt.Println(num)
	number, ok := bnlmz.Add(num[1], num[2])

	if ok != nil {
		fmt.Println(ok)
		return
	}

	fmt.Println(number)

	number, ok = bnlmz.Sub(num[1], num[2])

	if ok != nil {
		fmt.Println(ok)
		return
	}

	fmt.Println(number)
}
