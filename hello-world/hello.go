package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("****** String To Integer Convertion *****")
	stringToNum("hi")
}

func stringToNum(str string) (num int64, err error) {
	num, err = strconv.ParseInt(str, 10, 64)
	return
}
