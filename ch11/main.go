package main

import "fmt"

type number interface {
	count() float64
}

type double float64

func (n *double) count() float64 {
	return 0
}

func main() {
	var i double
	fmt.Println(i.count())
}
