package main

import (
	"fmt"
	"log"
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
		log.Fatalln("cannot parse string:", err)
		panic(err)
	}
	fmt.Println(num)
}
