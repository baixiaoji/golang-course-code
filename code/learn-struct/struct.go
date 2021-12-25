package main

import (
	"fmt"
	"os"
	"strconv"
)

type Product struct {
	id int
	name string
	description string
	price float64
}

func (prod *Product) PrintData()  {
	fmt.Printf("the product name: %v\n", prod.name)
	fmt.Printf("description: %v \n", prod.description)
	fmt.Printf("price: USD %.2f \n", prod.price)
}

func (prod Product) saveData()  {
	file, _ := os.Create(strconv.Itoa(prod.id) + ".txt")
	content := fmt.Sprint( fmt.Sprintf(
		"ID: %v\nTitle: %v\nDescription: %v\nPrice: USD %.2f\n",
		prod.id,
		prod.name,
		prod.description,
		prod.price,
	))
	file.WriteString(content)
	file.Close()
}
func main()  {
	firstProduce := getProduct(12,"Cars", "intell car easy for play", 23.3)
	firstProduce.PrintData()
	firstProduce.saveData()
}

func getProduct(id int, name string, description string, price float64) *Product {
	return &Product{id, name, description, price}
}