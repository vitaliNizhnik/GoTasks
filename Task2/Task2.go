package main

import "fmt"

type product struct{
	maker string
	typeOfProduct string
	categoryOfProduct string
	price float64
}
type laptop struct {
	product
	model string
	screenResolution string
	ram int
	hdmi bool
}

func(someProduct product) showProductInfo(){
	fmt.Println("This product was made by ",someProduct.maker,
		"\nType of the product: ",someProduct.typeOfProduct,
		"\nCategory : ",someProduct.categoryOfProduct,
		"\nPrice : ",someProduct.price)
}

func initArray()[4]product{
	var firstProduct = product{
		"Lenovo",
		"Electronic engenering",
		"PC",
		100}
	var secondProduct = product{
		"CoolMaker",
		"Computer",
		"PC",
		150}
	var thirdProduct = product{
		"NormalMaker",
		"Printer",
		"JetPrinter",
		150}
	var fourthProduct = product{
		"BadMaker",
		"Printer",
		"3d Printer",
		150}
	arr:= [4]product{firstProduct,secondProduct,thirdProduct,fourthProduct}
    return arr
}

func makeDiscount(someProduct *product,discountPercent float64){
	someProduct.price = discountPercent*someProduct.price
}

func showArray(arr *[4]product) {
	for i,num:=range arr{
		fmt.Printf("index: %v | Maker:%v; Category:%v; Type:%v; Price:%.1f\n",
			i,num.maker,num.categoryOfProduct,num.typeOfProduct,num.price)
	}
}

func main(){
 arr:=initArray()
 showArray(&arr)

 fmt.Scanf(" ")
}