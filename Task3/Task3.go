package main

import "fmt"

//Некий абстрактный продукт
type interProduct interface {
    print()
	makeDiscount(float64)
}
//Метод для вывода информации о product
func(someProduct product) print(){
	fmt.Println(" This is product made by ",someProduct.maker,
		"Type of the product: ",someProduct.typeOfProduct,
		"Category : ",someProduct.categoryOfProduct,
		"Price : ",someProduct.price)
}
//Метод для вывода информации о laptop
func(someProduct laptop) print(){
	fmt.Println(" This is Laptop model:",someProduct.model,
		"ScreenResolution: ",someProduct.screenResolution,
		"HDMI: ",someProduct.hdmi,
		"Price : ",someProduct.price)
}
//Метод для получения скидки у product
func(someProduct *product)makeDiscount(discountPercent float64){
	someProduct.price = discountPercent*someProduct.price
	fmt.Printf("Congratulations, you have %.2f discount for your Product, enjoy! ",discountPercent)
}
//Метод для получения скидки у laptop
func(someProduct *laptop)makeDiscount(discountPercent float64){
	someProduct.price = discountPercent*someProduct.price
	fmt.Printf("Congratulations, you have %.2f discount for your Laptop, enjoy! ",discountPercent)
}
// Инициализация массива как элементами product, так и laptop
func initArray()[4]interProduct{

	var firstProduct = product{
		"MSI",
		"Electronic",
		"PC",
		100}

	var secondProduct  = laptop{ product{},
		"K640",
		"1920x1080",
		500,
		false,
		250}

	var thirdProduct = product{
		"NormalMaker",
		"Printer",
		"JetPrinter",
		150}

	var fourthProduct = laptop{ product{},
		"K640",
		"1920x1080",
		500,
		false,
		300}
	arr:= [4]interProduct{&firstProduct,&secondProduct,&thirdProduct,&fourthProduct}
	return arr
}
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
	price float64
}

func showArray(arr *[4]interProduct) {
	for i:=0; i<len(arr);i++{
		fmt.Print("Index :",i)
		arr[i].print()
	}
}

func main(){

	arr:= initArray()
	showArray(&arr)

	//Проверка работы метода makeDiscount для разных типов
	arr[0].makeDiscount(0.5)
	arr[0].print()
	arr[1].makeDiscount(0.5)
	arr[1].print()
	fmt.Scanf(" ")
}