package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i == 5 {
			continue
		} else if i == 10 {
			break
		}
		fmt.Println(i)
	}
}
