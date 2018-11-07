package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const BADREQUEST = -100 //константа для неудачных запросов

func main() {
	var url = flag.String("url", "https://www.google.com/", "help message for URL")
	var requestNumber = flag.Int("reqnum", 5, "This is a quantity of request's")
	var timeout = flag.Int("timeout", 500, "This is a maximum response time")

	//Клиент для создания запросов
	client := http.Client{
		Timeout: time.Duration(*timeout) * time.Millisecond,
	}

	flag.Parse()
	fmt.Println("url = ",*url)
	fmt.Println("request number = ",*requestNumber)
	fmt.Println("timeout = ",*timeout)

	//Слайс для хранения времени за которое отработал каждый запрос
	var timeArray = []int{}

	//Время начала работы запросов
	startTime:= time.Now()

	wg:= &sync.WaitGroup{}
	for i:=0; i<*requestNumber; i++{
		wg.Add(1)
		go MakeRequest(*url,wg,client,&timeArray)
	}
	wg.Wait()

	//Время окончания работы запросов
	finishTime:=time.Now()

	//Время за которое все запросы были выполнены
	resultTime := finishTime.Sub(startTime)

	fmt.Println("Time for all request's:",resultTime)
	analizeAndShow(timeArray,*timeout)

}
// Создаем запрос и получаем значение времени, которое затем записываем в слайс
func MakeRequest(url string, wg *sync.WaitGroup, client http.Client,timeArray *[]int ){
	defer wg.Done()

	var mutex = &sync.Mutex{}

	startTime:= time.Now()

	//Отправка запроса
	resp, err := client.Get(url)
	if err != nil {
		//Добавление к слайсу нового элемента в случае, если запрос не удачен
		// А именно, если запрос неуспешен, ему присваивается отрицательное время, что впринципе невозможно
		mutex.Lock()
		*timeArray = append(*timeArray,BADREQUEST)
		mutex.Unlock()

	}
	defer resp.Body.Close()

	finishTime:=time.Now()
	elapsed := finishTime.Sub(startTime)
	var resultTime = int(elapsed/time.Millisecond)

	runtime.Gosched()
	//Добавление к слайсу нового элемента
	mutex.Lock()
	*timeArray = append(*timeArray,resultTime)
	mutex.Unlock()
}
  //Обработка и вывод данных
func analizeAndShow(timeArray []int,timeout int){

	for i:=0; i<len(timeArray); i++{
		fmt.Println(i+1," request time: ",timeArray[i])
	}
	var min = findMinInSlice(timeArray,timeout)
	fmt.Println("Min response time:", min)

	var max = findMaxInSlice(timeArray)
	fmt.Println("Max response time:", max)

	var failed = findNumberOfFailedRequests(timeArray)
	fmt.Println("Number of failed request's time:", failed)

	var avg = findAverageRequestTime(timeArray)
	fmt.Println("Average response time:", avg)
	fmt.Scanf(" ")

}

func findMinInSlice(timeArray []int,timeout int)int{
	var min= timeout
	for i:=0; i<len(timeArray); i++{
		if timeArray[i] < min && timeArray[i]!=BADREQUEST{
			min = timeArray[i]
		}
	}
	return min
}

func findMaxInSlice(timeArray []int)int{
	var max= 0
	for i:=0; i<len(timeArray); i++{
		if timeArray[i] > max{
			max = timeArray[i]
		}
	}
	return max
}

func findAverageRequestTime(timeArray []int)int{
	var sum = 0
	for i:=0; i<len(timeArray); i++{
		if timeArray[i] != BADREQUEST {
			sum += timeArray[i]
		}
	}
	return sum/(len(timeArray)-findNumberOfFailedRequests(timeArray))
}

func findNumberOfFailedRequests(timeArray []int)int{
	var number = 0
	for i:=0; i<len(timeArray); i++{
		if timeArray[i] == BADREQUEST {
			number ++
		}
	}
	return number
}