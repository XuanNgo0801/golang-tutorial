package main

import (
	"fmt"
)

func hello(done chan bool) { // Tao ra 1 channel bool va truyen duoi dang tham so vao Goroutine hello
	fmt.Println("Hello world goroutine")
	done <- true
}
func main() {
	done := make(chan bool) // tao 1 channeldone co kieu bool
	go hello(done)
	<-done
	fmt.Println("main function")
}
