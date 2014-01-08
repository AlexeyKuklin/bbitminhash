package main

import (
	"fmt"
)

func main() {
	res := Minhash([]string{
		"aaa", "bbb", "ccc", "ddd", "eee",
	}, []string{
		"aaa", "bbb", "ccc", "ddd",
	})

	fmt.Println("res = ", res)
}
