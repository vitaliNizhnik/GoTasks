package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

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
	ShowMetrix(timeArray,*timeout,*requestNumber)
}

// Создаем запрос и получаем значение времени, которое затем записываем в слайс
func MakeRequest(url string, wg *sync.WaitGroup, client http.Client,timeArray *[]int ){
	defer wg.Done()

	var mutex = &sync.Mutex{}

	startTime:= time.Now()

	//Отправка запроса
	resp, err := client.Get(url)
	if err != nil {
		return

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

func getMetrix(timeArray []int,timeout int,requestNumber int)(int,int,int,int){

	var min= timeout
	var max = 0
	var sum = 0
	var failedNumber = requestNumber - len(timeArray)
	for i:=0; i<len(timeArray); i++{
		if timeArray[i] < min{
			min = timeArray[i]
		}
		if timeArray[i] > max{
			max = timeArray[i]
		}
		sum+=timeArray[i]
	}
	return min, max, sum/(len(timeArray)),failedNumber
}
//Обработка и вывод данных
func ShowMetrix(timeArray []int,timeout int, requestNumber int){

	for i:=0; i<len(timeArray); i++{
		fmt.Println(i+1," request time: ",timeArray[i])
	}

	var min, max, avg, failed = getMetrix(timeArray,timeout,requestNumber)

	fmt.Println("Min response time:", min)
	fmt.Println("Max response time:", max)
	fmt.Println("Number of failed request's time:", failed)
	fmt.Println("Average response time:", avg)
	fmt.Scanf(" ")
}