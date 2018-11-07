package main

import (
	"fmt"
	"sync"
)

func main(){
	//задание 4
	func(){
	defer fmt.Println("World")
		  fmt.Println("Hello")
	}()
	//задание 4.1**
	once:= sync.Once{}
	for i:=0; i<10; i++{
		once.Do(sayHello)
	}
	fmt.Scanf(" ")
}

func sayHello(){
	defer fmt.Println("World")
	      fmt.Println("Hello")
}