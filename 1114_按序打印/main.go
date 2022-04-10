package main

import (
	"fmt"
	"time"
)

var oneDone = make(chan int)
var twoDone = make(chan int)

func first() {
	fmt.Println("1")
	oneDone <- 1
}
func seccond() {
	<-oneDone
	fmt.Println("2")
	twoDone <- 1
}
func thrid() {
	<-twoDone
	fmt.Println("3")
}
func main() {
	go first()
	go seccond()
	go thrid()
	time.Sleep(time.Second) //防止主线程退出，子线程没运行完
}
