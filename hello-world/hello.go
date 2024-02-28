package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("****** String To Integer Convertion *****")
	stringToNum()
}

func stringToNum() {
	str := "hi"
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}
