package main

import "fmt"

//顺序打印

func main() {
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)
	c3 := make(chan struct{}, 1)
	b := make(chan struct{}) //阻塞主线程
	c1 <- struct{}{}
	go dog(c1, c2)
	go fish(c2, c3)
	go cat(c3, c1, b)
	b <- struct{}{}
}

func dog(c1, c2 chan struct{}) {
	for i := 0; i < 10; i++ {
		<-c1
		fmt.Println("dog", i)
		c2 <- struct{}{}
	}
}

func fish(c2, c3 chan struct{}) {
	for i := 0; i < 10; i++ {
		<-c2
		fmt.Println("fish", i)
		c3 <- struct{}{}
	}
}

func cat(c3, c1, b chan struct{}) {
	for i := 0; i < 10; i++ {
		<-c3
		fmt.Println("cat", i)
		c1 <- struct{}{}
	}
	<-b
}
