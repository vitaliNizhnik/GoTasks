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

	wg:= sync.WaitGroup{}

	//задание 4.1**
	sayHello := func(){
		defer fmt.Println("World")
		fmt.Println("Hello")

	}

	wg.Add(10)
	once:= sync.Once{}
	for i:=0; i<10; i++{
		go func(){
			once.Do(sayHello)
		}()
		wg.Done()
	}
	wg.Wait()


}
