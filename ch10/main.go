package main

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func main() {
	fmt.Println(f1())
}
