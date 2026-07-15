package main

import (
	"fmt"
	"learngo/hello"
)

func main() {
	message := hello.Welcome("Daniel")
	fmt.Println(message)
}
