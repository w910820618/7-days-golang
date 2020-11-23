package main

import "fmt"

type People interface {
	Read() string
}

type Man struct {
	name string
}

func (m Man) Read() string {
	return "Hello"
}

func main() {
	var people People
	people = Man{}
	fmt.Println(people.Read())
}
