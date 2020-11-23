package main

import (
	"fmt"
	"strconv"
)

type saveLog func(msg string)

func stringToInt(s string, log saveLog) int64 {

	if value, err := strconv.ParseInt(s, 0, 0); err != nil {
		log(err.Error())
		return 0
	} else {
		return value
	}
}

func myLog(msg string) {
	fmt.Println("Find Error:", msg)
}

func main() {
	stringToInt("123", myLog)
}
