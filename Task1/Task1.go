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
func makeDiscount(someProduct *product,discountPercent float64){
	someProduct.price = discountPercent*someProduct.price
}
func main(){
	var firstProduct = product{
		"Lenovo",
		"Electronic engenering",
		"PC",
		100}
	var secondProduct = laptop{
		product{
		"MSI",
		"Electronic engenering",
		"Laptop",
		200},
		"Z510",
		"1920x1080",
		4096,
		true }

	firstProduct.showProductInfo()
	makeDiscount(&firstProduct,0.5)

	fmt.Println("\nProduct info after discount")
	firstProduct.showProductInfo()

	fmt.Println("\nInner struct info")
	secondProduct.showProductInfo()

	fmt.Scanf(" ")
}
